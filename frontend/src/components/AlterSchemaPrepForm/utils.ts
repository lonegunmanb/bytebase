import axios from "axios";
import { useSilentRequest } from "@/plugins/silent-request";
import { useSchemaEditorV1Store } from "@/store";
import { DatabaseEdit, DatabaseEditResult } from "@/types";
import {
  checkHasSchemaChanges,
  diffSchema,
  mergeDiffResults,
} from "@/utils/schemaEditor/diffSchema";

export const getDatabaseEditListWithSchemaEditor = () => {
  const schemaEditorV1Store = useSchemaEditorV1Store();

  const databaseEditList: DatabaseEdit[] = [];
  for (const databaseSchema of Array.from(
    schemaEditorV1Store.resourceMap["database"].values()
  )) {
    const database = databaseSchema.database;
    for (const schema of databaseSchema.schemaList) {
      const originSchema = databaseSchema.originSchemaList.find(
        (originSchema) => originSchema.id === schema.id
      );
      if (!originSchema) {
        continue;
      }

      const diffSchemaResult = diffSchema(database.name, originSchema, schema);
      if (checkHasSchemaChanges(diffSchemaResult)) {
        const index = databaseEditList.findIndex(
          (edit) => String(edit.databaseId) === database.uid
        );
        if (index !== -1) {
          databaseEditList[index] = {
            databaseId: Number(database.uid),
            ...mergeDiffResults([diffSchemaResult, databaseEditList[index]]),
          };
        } else {
          databaseEditList.push({
            databaseId: Number(database.uid),
            ...diffSchemaResult,
          });
        }
      }
    }
  }
  return databaseEditList;
};

export const postDatabaseEdit = async (
  databaseEdit: DatabaseEdit,
  silent = false
) => {
  const request = () => {
    const post = () =>
      axios.post(`/api/database/${databaseEdit.databaseId}/edit`, databaseEdit);
    if (silent) {
      return useSilentRequest(post);
    }
    return post();
  };

  const resData = (await request()).data;
  const databaseEditResult = resData.data.attributes as DatabaseEditResult;
  return databaseEditResult;
};
