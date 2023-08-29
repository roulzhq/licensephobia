"use client";

import { useMemo, useState } from "react";
import { getAllLicenses } from "core";

import styles from "./page.module.scss";

const allLicenses = getAllLicenses().sort(
  (a, b) => a.name?.localeCompare(b.name)
);

export default function Page(): JSX.Element {
  const [filterString, setFilterString] = useState<string>("");
  const [includeRareLicenses, setIncludeRareLicenses] =
    useState<boolean>(false);

  const filteredLicenses = useMemo(() => {
    const lowercaseSearchString = filterString.toLowerCase();

    return allLicenses.filter((i) => {
      if (!includeRareLicenses && !i.calFound) return false;

      return (
        i.name?.toLowerCase().includes(lowercaseSearchString) ||
        i.description?.toLowerCase().includes(lowercaseSearchString)
      );
    });
  }, [filterString, includeRareLicenses]);

  return (
    <div className={styles.page}>
      <div className={styles.header}>
        <h1>Licenses</h1>
        <div>
          <input
            type="search"
            placeholder="Search licenses"
            value={filterString}
            onChange={(e) => setFilterString(e.target.value)}
          />
          <label>
            Show rare licenses
            <input
              type="checkbox"
              checked={includeRareLicenses}
              onChange={() => setIncludeRareLicenses((c) => !c)}
            />
          </label>
        </div>
      </div>
      <div className={styles.licenseList}>
        <ul>
          {filteredLicenses.map((i) => (
            <li key={i.id}>
              <b>{i.name}</b>
              <div>
                <p>{i.description ?? "No description known"}</p>
                <a href={i.spdxReferenceUrl} target="_blank">{i.spdxReferenceUrl}</a>
              </div>
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
}
