/**
 * This script fetches SPDX License metadata from their github repository at
 * https://github.com/spdx/license-list-data
 * 
 * The data is then stored unchanged in a JSON file.
 */

import * as fs from "node:fs/promises";

import { SPDXResponse } from "./types";

const SPDX_LICENSE_URL =
  "https://raw.githubusercontent.com/spdx/license-list-data/master/json/licenses.json";

const OUTPUT_JSON_PATH = "./public/spdx.json";

async function fetchSpdxLicenseJson(): Promise<SPDXResponse | null> {
  try {
    const res = await fetch(encodeURI(SPDX_LICENSE_URL));
    const json = await res.json();

    return json;
  } catch (error) {
    console.error("Error while downloading SPDX data", error);
  }

  return null;
}

async function fetchAndSave() {
  try {
    const json = await fetchSpdxLicenseJson();

    if (json)
      await fs.writeFile(OUTPUT_JSON_PATH, JSON.stringify(json.licenses, null, 2));
  } catch (error) {
    console.error("Error while writing JSON to file", error);
  }
}

fetchAndSave();
