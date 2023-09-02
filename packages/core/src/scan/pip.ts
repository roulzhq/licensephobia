import { PackageScanResult, PipfileInput } from "@licensephobia/types";
import toml from "toml";

const SUPPORTED_PIPFILE_DEPENDENCY_FIELDS = [
  "dev-packages",
  "packages",
] as const;

export function isPipfile(pipObj: Record<string, any>): pipObj is PipfileInput {
  const includesAndIsObject = (key: string) => {
    if (key in pipObj) {
      return typeof pipObj[key] === "object";
    }

    return false;
  };

  return SUPPORTED_PIPFILE_DEPENDENCY_FIELDS.some(includesAndIsObject);
}

export function scanPipfile(pipfileString: string): PackageScanResult[] {
  const pipfile = toml.parse(pipfileString);

  if (!isPipfile(pipfile)) {
    throw new Error("Provided TOML is not a valid Pipfile");
  }

  const results = SUPPORTED_PIPFILE_DEPENDENCY_FIELDS.flatMap((objName) => {
    if (!(objName in pipfile)) return [];
    return Object.entries(pipfile[objName])
      .map(([name, version]) => {
        if (typeof version !== "string") return undefined;

        return {
          name,
          source: objName,
          version,
          licenseId: null,
          websiteUrl: null,
        };
      })
      .filter((i) => i !== undefined) as PackageScanResult[];
  });

  return results;
}
