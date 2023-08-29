"use client"

import SearchField from "$components/searchField";
import styles from "./page.module.scss";

export default function Page(): JSX.Element {
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
      </div>
    </div>
  );
}
