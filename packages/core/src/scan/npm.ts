import {
  DependencyFileScanResult,
  PackageJson,
  PackageScanResult,
} from "@licensephobia/types";

const SUPPORTED_PACKAGE_JSON_DEPENDENCY_FIELDS = [
  "dependencies",
  "devDependencies",
  "peerDependencies",
] as const;

export function isPackageJson(
  jsonObj: Record<string, unknown>
): jsonObj is PackageJson {
  const includesAndIsObject = (key: string) => {
    if (key in jsonObj) {
      return typeof jsonObj[key] === "object";
    }

    return false;
  };

  return SUPPORTED_PACKAGE_JSON_DEPENDENCY_FIELDS.some(includesAndIsObject);
}

export function scanPackageJson(jsonString: string): DependencyFileScanResult {
  const packageJson = JSON.parse(jsonString);

  if (!isPackageJson(packageJson)) {
    throw new Error("Provided JSON is not a valid package.json");
  }

  const results: PackageScanResult[] =
    SUPPORTED_PACKAGE_JSON_DEPENDENCY_FIELDS.flatMap((objName) => {
      if (!(objName in packageJson)) return [];

      return Object.entries(packageJson[objName]).map(([id, version]) => {
        if(typeof version !== "string") throw new Error("Unexpected error parsing package.lock")

        return {
          id,
          source: objName,
          version,
        };
      });
    });

  return {
    packageManager: "npm",
    results,
  };
}