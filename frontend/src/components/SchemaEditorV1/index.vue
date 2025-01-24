<template>
  <div
    v-if="state.initialized"
    class="w-full h-full border rounded-lg overflow-hidden"
    v-bind="$attrs"
  >
    <Splitpanes
      class="default-theme w-full h-full flex flex-row overflow-hidden"
    >
      <Pane min-size="15" size="25">
        <Aside />
      </Pane>
      <Pane min-size="60" size="75">
        <Editor />
      </Pane>
    </Splitpanes>
  </div>
</template>

<script lang="ts" setup>
import { isEqual } from "lodash-es";
import { Splitpanes, Pane } from "splitpanes";
import { onMounted, watch, reactive } from "vue";
import { useSchemaEditorV1Store, useSettingV1Store } from "@/store";
import { useSchemaDesignStore } from "@/store/modules/schemaDesign";
import { ComposedProject, ComposedDatabase } from "@/types";
import {
  SchemaDesign,
  SchemaDesign_Type,
} from "@/types/proto/v1/schema_design_service";
import Aside from "./Aside/index.vue";
import Editor from "./Editor.vue";
import { convertBranchToBranchSchema } from "./utils/branch";

const props = defineProps<{
  project: ComposedProject;
  resourceType: "database" | "branch";
  readonly?: boolean;
  databases?: ComposedDatabase[];
  // NOTE: we only support editing one branch for now.
  branches?: SchemaDesign[];
}>();

interface LocalState {
  initialized: boolean;
}

const settingStore = useSettingV1Store();
const schemaEditorV1Store = useSchemaEditorV1Store();
const schemaDesignStore = useSchemaDesignStore();
const state = reactive<LocalState>({
  initialized: false,
});

const prepareBranchContext = async () => {
  if (props.resourceType !== "branch" || !props.branches) {
    return;
  }
  for (const branch of props.branches) {
    if (branch.type === SchemaDesign_Type.PERSONAL_DRAFT) {
      // Prepare parent branch for personal draft.
      await schemaDesignStore.getOrFetchSchemaDesignByName(
        branch.baselineSheetName
      );
    }
  }
};

const initialSchemaEditorState = () => {
  schemaEditorV1Store.setState({
    project: props.project,
    resourceType: props.resourceType,
    // NOTE: this will clear all tabs. We will restore tabs as needed later.
    tabState: {
      tabMap: new Map(),
    },
    readonly: props.readonly || false,
  });

  if (props.resourceType === "database") {
    schemaEditorV1Store.setState({
      resourceMap: {
        // NOTE: we will dynamically fetch schema list for each database in database tree view.
        database: new Map(
          (props.databases || []).map((database) => [
            database.name,
            {
              database,
              schemaList: [],
              originSchemaList: [],
            },
          ])
        ),
        branch: new Map(),
      },
    });
  } else {
    schemaEditorV1Store.setState({
      resourceMap: {
        database: new Map(),
        branch: new Map(
          (props.branches || []).map((branch) => [
            branch.name,
            convertBranchToBranchSchema(branch),
          ])
        ),
      },
    });
  }
};

// Prepare schema template contexts.
onMounted(async () => {
  await settingStore.getOrFetchSettingByName("bb.workspace.schema-template");
  await prepareBranchContext();

  initialSchemaEditorState();
  state.initialized = true;
});

watch(
  () => props,
  (newProps, oldProps) => {
    schemaEditorV1Store.setState({
      project: newProps.project,
      resourceType: newProps.resourceType,
      readonly: newProps.readonly || false,
    });

    // Update editor state if needed.
    // * If we update databases/branches, we need to rebuild the editing state.
    // * If we update databases/branches, we need to clear all tabs.
    if (newProps.resourceType === "database") {
      if (isEqual(newProps.databases, oldProps.databases)) {
        return;
      }
      schemaEditorV1Store.setState({
        tabState: {
          tabMap: new Map(),
        },
        resourceMap: {
          // NOTE: we will dynamically fetch schema list for each database in database tree view.
          database: new Map(
            (props.databases || []).map((database) => [
              database.name,
              {
                database,
                schemaList: [],
                originSchemaList: [],
              },
            ])
          ),
          branch: new Map(),
        },
      });
    } else {
      if (isEqual(newProps.branches, oldProps.branches)) {
        return;
      }
      schemaEditorV1Store.setState({
        tabState: {
          tabMap: new Map(),
        },
        resourceMap: {
          database: new Map(),
          branch: new Map(
            (props.branches || []).map((branch) => [
              branch.name,
              convertBranchToBranchSchema(branch),
            ])
          ),
        },
      });
    }
  },
  {
    deep: true,
  }
);
</script>

<style>
@import "splitpanes/dist/splitpanes.css";

/* splitpanes pane style */
.splitpanes.default-theme .splitpanes__pane {
  @apply bg-transparent !transition-none;
}

.splitpanes.default-theme .splitpanes__splitter {
  @apply bg-gray-100 border-none;
}

.splitpanes.default-theme .splitpanes__splitter:hover {
  @apply bg-indigo-300;
}

.splitpanes.default-theme .splitpanes__splitter::before,
.splitpanes.default-theme .splitpanes__splitter::after {
  @apply bg-gray-700 opacity-50 text-white;
}

.splitpanes.default-theme .splitpanes__splitter:hover::before,
.splitpanes.default-theme .splitpanes__splitter:hover::after {
  @apply bg-white opacity-100;
}
</style>
