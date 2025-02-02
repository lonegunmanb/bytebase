import { defaultTo } from "lodash-es";
import {
  AddColumnContext,
  CreateTableContext,
  ChangeColumnContext,
} from "@/types";
import { Column, Table } from "@/types/v1/schemaEditor";

export const getColumnComment = (column: Column): string => {
  return [
    defaultTo(column.classification, ""),
    defaultTo(column.userComment, ""),
  ]
    .filter((val) => val)
    .join("-");
};

export const transformColumnToAddColumnContext = (
  column: Column
): AddColumnContext => {
  return {
    name: defaultTo(column.name, ""),
    type: defaultTo(column.type, ""),
    comment: getColumnComment(column),
    nullable: defaultTo(column.nullable, false),
    default: defaultTo(column.default, undefined),
    characterSet: "",
    collation: "",
  };
};

export const transformColumnToChangeColumnContext = (
  originColumn: Column,
  column: Column
): ChangeColumnContext => {
  return {
    oldName: defaultTo(originColumn.name, ""),
    newName: defaultTo(column.name, ""),
    type: defaultTo(column.type, ""),
    comment: getColumnComment(column),
    nullable: defaultTo(column.nullable, false),
    default: defaultTo(column.default, undefined),
    characterSet: "",
    collation: "",
  };
};

export const transformTableToCreateTableContext = (
  schema: string,
  table: Table
): CreateTableContext => {
  return {
    schema,
    name: defaultTo(table.name, ""),
    engine: defaultTo(table.engine, ""),
    collation: defaultTo(table.collation, ""),
    comment: defaultTo(table.comment, ""),
    addColumnList: [],
    primaryKeyList: [],
    addForeignKeyList: [],
    // As we don't have a CharacterSet field in table model,
    // set it as an empty string for now.
    characterSet: "",
  };
};
