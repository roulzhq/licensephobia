"use client";

import { scanPackageJson } from "@licensephobia/core";
import { useRouter } from "next/navigation";
import styles from "./page.module.scss";
import Dropzone from "$components/dropzone";
import SearchField from "$components/searchField";

export default function Page(): JSX.Element {
  const router = useRouter();

  const handleFileSubmit = (file: File) => {
    // eslint-disable-next-line @typescript-eslint/no-unsafe-assignment, @typescript-eslint/no-unsafe-call
    const reader = new FileReader();
    // eslint-disable-next-line @typescript-eslint/no-unsafe-member-access, @typescript-eslint/no-unsafe-call
    reader.readAsText(file);
    // eslint-disable-next-line @typescript-eslint/no-unsafe-member-access, @typescript-eslint/no-unsafe-call
    reader.onload = (e) => {
      // eslint-disable-next-line @typescript-eslint/no-unsafe-member-access
      if (e?.target?.result) {
        // eslint-disable-next-line @typescript-eslint/no-unsafe-member-access
        const data = e.target.result as string;

        try {
          const scanResult = scanPackageJson(data);

          const redirectUrl = `/scan?${scanResult.results
            .map((i) => `${i.name}@${i.version}`)
            .join("&package=")}`;
          router.push(redirectUrl);
        } catch (err) {
          // eslint-disable-next-line no-console
          console.log(err);
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
          onBlur={console.log}
          onChange={console.log}
          onHintClick={console.log}
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
