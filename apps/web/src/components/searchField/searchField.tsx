"use client"

import { ChangeEvent, ChangeEventHandler, MouseEventHandler } from "react";

import styles from "./searchField.module.scss";

export interface SearchFieldProps {
  placeholder: string;
  hints: string[];
  onChange: ChangeEventHandler<HTMLInputElement>;
  onBlur: MouseEventHandler<HTMLButtonElement>;
  onHintClick: (e: Event) => void;
}

export default function SearchField({
  placeholder,
  hints,
  onChange,
  onBlur,
  onHintClick,
}: SearchFieldProps) {
  return (
    <div className={styles.container}>
      <input type="search" onChange={onChange} />
      <button type="submit" onClick={onBlur}>
        Search
      </button>
    </div>
  );
}
