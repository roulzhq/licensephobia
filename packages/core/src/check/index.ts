import { PackageManagerId } from "@licensephobia/types";
import { checkNpmPackages } from "./npm";

export * from "./npm";

export async function checkPackages(
  packageManagerId: PackageManagerId,
  packages: [string, string][]
) {
  switch (packageManagerId) {
    case "npm": {
      return checkNpmPackages(packages);
    }
    default: {
      throw new Error("Unknown package manager. Could not check packages.");
    }
  }
}
