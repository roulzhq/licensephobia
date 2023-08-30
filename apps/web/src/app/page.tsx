"use client";

import { useRouter } from "next/navigation";
import { scanPackageJson } from "core";

import Dropzone from "$components/dropzone";
import SearchField from "$components/searchField";
import styles from "./page.module.scss";

export default function Page(): JSX.Element {
  const router = useRouter();

  const handleFileSubmit = (file: File) => {
    const reader = new FileReader();
    reader.readAsText(file);
    reader.onload = (e) => {
      if (e?.target?.result) {
        const data = e.target.result as string;

        try {
          const scanResult = scanPackageJson(data);

          const redirectUrl = `/scan?${scanResult.results
            .map((i) => `${i.name}@${i.version}`)
            .join("&package=")}`;
          router.push(redirectUrl);
        } catch (e) {
          console.log(e);
        }
      }
    };
  };

  return (
    <div className={styles.page}>
      <h1 className={styles.title}>
        Don&apos;t be afraid of{" "}
        <span className={styles.titleChip}>Node.js</span> Licenses anymore!
      </h1>
      <div>
        <SearchField
          placeholder="search"
          onBlur={() => {}}
          onChange={() => {}}
          onHintClick={() => {}}
          hints={[]}
        />
        <Dropzone
          placeholder="upload"
          onChange={console.log}
          onSubmit={handleFileSubmit}
        />
      </div>
    </div>
  );
}
