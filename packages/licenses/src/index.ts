import SPDXJson from "../public/spdx.json";
import CALJson from "../public/cal.json";
import { CALLicenseDefinition, SPDXLicenseDefinition } from "./types";

const spdxData = SPDXJson as SPDXLicenseDefinition[];
const calData = CALJson as CALLicenseDefinition[];

export { spdxData, calData };
