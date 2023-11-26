import ky from "ky";
import sv from "semver";

import { NpmRegPackage, PackageCheckResult } from "@licensephobia/types";

const NPM_REG_BASE_URL = "https://registry.npmjs.org";

export async function checkNpmPackage(
  id: string,
  version: string
): Promise<PackageCheckResult | null> {
  const versionRange = new sv.Range(version);
  const minVersionFromRange = sv.minVersion(versionRange)?.version;

  if (!sv.valid(minVersionFromRange)) {
    throw new Error(
      `Version identifier ${version} for package ${id} is invalid.`
    );
  }

  try {
    const res = (await ky
      .get(`${NPM_REG_BASE_URL}/${id}/${minVersionFromRange}`, { retry: 0 })
      .json()) as NpmRegPackage;

    const { _id, name, description, license, homepage } = res;

    return {
      id: _id,
      name,
      description,
      licenseId: license,
      homepage,
    };
  } catch (e) {
    console.log(e);
    return null;
  }
}

export async function checkNpmPackages(
  packages: [string, string][]
): Promise<Record<string, PackageCheckResult>> {
  const results = await Promise.allSettled(
    packages.map((i) => checkNpmPackage(i[0], i[1]))
  );

  return results.reduce(
    (acc, curr) => {
      if (curr.status === "fulfilled") {
        const { value } = curr as PromiseFulfilledResult<PackageCheckResult>;

        acc[value.id] = value;
        return acc;
      }

      return acc;
    },
    {} as Record<string, PackageCheckResult>
  );
}
