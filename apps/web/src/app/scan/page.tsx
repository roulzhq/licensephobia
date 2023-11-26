"use client";

import { ReadonlyURLSearchParams, useSearchParams } from "next/navigation";
import Error from "next/error";

import { scanFileString } from "@licensephobia/core";
import {
  DependencyFileScanResult,
  PackageManagerId,
} from "@licensephobia/types";
import { useCheckPackages } from "$api/queries/packages";

function getPackagesFromUrl(
  searchParams: ReadonlyURLSearchParams
): DependencyFileScanResult | null {
  const packageManager = searchParams.get("pm") as PackageManagerId;
  const base64FileString = searchParams.get("file");

  if (packageManager == null || base64FileString == null) {
    return null;
  }

  const base64String = decodeURIComponent(base64FileString);
  const fileBuffer = Buffer.from(base64String, "base64");
  const fileString = fileBuffer.toString("utf8");

  let scanResult: DependencyFileScanResult | null;

  try {
    scanResult = scanFileString(packageManager, fileString);
  } catch (e) {
    return null;
  }

  return scanResult;
}

export default function Page(): JSX.Element {
  const searchParams = useSearchParams();
  const fileScanResult = getPackagesFromUrl(searchParams);
  const { data: checkedPackages } = useCheckPackages(
    fileScanResult?.packageManager!,
    fileScanResult?.results.map((i) => [i.id, i.version]) ?? []
  );

  if (fileScanResult == null) return <Error statusCode={500} />;

  return (
    <div>
      <h1>Scanning</h1>
      <pre>{JSON.stringify(checkedPackages, null, 4)}</pre>
    </div>
  );
}
