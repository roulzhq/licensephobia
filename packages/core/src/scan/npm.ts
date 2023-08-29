import { DependencyFileScanResult } from "types";

export function scanPackageJson(jsonBlob: any): DependencyFileScanResult {
  return {
    packageManager: "npm",
    results: [],
  };
}
