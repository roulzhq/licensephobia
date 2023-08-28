/**
 * This script fetches License permissions, conditions and limitations
 * from the Choosealicense.com Github repository that can be found at
 * https://github.com/github/choosealicense.com
 *
 * The data is then converted to JSON and stored in a file.
 *
 * The script uses the SPDX license data fetched to choose which licenses to fetch.
 */

import * as fs from "node:fs/promises";
import { parseAllDocuments } from "yaml";

import { CALLicenseDefinition } from "./types";
import { spdxData } from ".";

const CAL_LICENSE_BASE_PATH =
  "https://raw.githubusercontent.com/github/choosealicense.com/gh-pages/_licenses";

const OUTPUT_JSON_PATH = "./public/cal.json";

async function fetchLicense(
  spdxId: string
): Promise<CALLicenseDefinition | null> {
  const url = `${CAL_LICENSE_BASE_PATH}/${spdxId.toLowerCase()}.txt`;

  try {
    const res = await fetch(encodeURI(url));
    if (res.ok) {
      const text = await res.text();
      const yaml = parseAllDocuments(text);

      return yaml[0].toJSON() as CALLicenseDefinition;
    }
  } catch (error) {
    console.error(
      "Error while downloading Choosealicense.com data for",
      spdxId,
      error
    );
  }

  return null;
}

async function fetchAndSave() {
  const promises = spdxData.flatMap((i) => {
    const res = fetchLicense(i.licenseId);

    return res ?? [];
  });

  const promiseResult = await Promise.allSettled(promises);
  const licenses = promiseResult
    .flatMap((i) => (i.status === "fulfilled" ? i.value ?? [] : []))
    .sort((a, b) => a["spdx-id"].localeCompare(b["spdx-id"]));

  try {
    if (licenses)
      await fs.writeFile(OUTPUT_JSON_PATH, JSON.stringify(licenses, null, 2));
  } catch (error) {
    console.error("Error while writing JSON to file", error);
  }
}

fetchAndSave();
