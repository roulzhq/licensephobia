"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";

import { PackageManagerId } from "@licensephobia/types";

import Dropzone from "$components/dropzone";
import { fileToBase64 } from "$utils";

import styles from "./page.module.scss";

export default function Page(): JSX.Element {
  const router = useRouter();
  const [packageManager] = useState<PackageManagerId>("npm");

  async function handleFileSubmit(file: File) {
    try {
      const fileBase64 = (await fileToBase64(file)).slice(
        "data:application/json;base64,".length
      );
      const redirectUrl = `/scan?pm=${packageManager}&file=${encodeURIComponent(
        fileBase64
      )}`;
      router.push(redirectUrl);
    } catch (e) {
      console.log("error parsing file", e);
    }
  }

  return (
    <div className={styles.page}>
      <h1 className={styles.title}>
        Don&apos;t be afraid of{" "}
        <span className={styles.titleChip}>Node.js</span> Licenses anymore!
      </h1>
      <div>
        <Dropzone
          placeholder="upload"
          onChange={console.log}
          onSubmit={handleFileSubmit}
        />
      </div>
    </div>
  );
}
