import { PackageManagerId } from "@licensephobia/types";
import { scanPackageJson } from "./npm";

export * from "./npm";

export function scanFileString(packageManagerId: PackageManagerId, fileString: string) {
  switch (packageManagerId) {
    case "npm": {
      return scanPackageJson(fileString);
    }
    default: {
      throw new Error("Unknown package manager. Could not scan file.");
    }
  }
}
