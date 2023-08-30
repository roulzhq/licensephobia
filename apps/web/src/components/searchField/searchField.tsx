"use client";

import type { ChangeEventHandler, MouseEventHandler } from "react";
import styles from "./searchField.module.scss";

export interface SearchFieldProps {
  placeholder: string;
  hints: string[];
  onChange: ChangeEventHandler<HTMLInputElement>;
  onBlur: MouseEventHandler<HTMLButtonElement>;
  onHintClick: (e: Event) => void;
}

export default function SearchField({ onChange, onBlur }: SearchFieldProps): JSX.Element {
  return (
    <div className={styles.container}>
      <input type="search" onChange={onChange} />
      <button type="submit" onClick={onBlur}>
        Search
      </button>
    </div>
  );
}
