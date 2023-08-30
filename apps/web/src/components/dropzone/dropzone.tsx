"use client";

import type { ChangeEventHandler, MouseEvent, MouseEventHandler } from "react";
import { useState } from "react";
import styles from "./dropzone.module.scss";

export interface DropzoneProps {
  placeholder: string;
  onChange?: ChangeEventHandler<HTMLInputElement>;
  onSubmit?: (file: File, e: MouseEvent<HTMLButtonElement>) => void;
}

export default function Dropzone({
  onChange,
  onSubmit,
}: DropzoneProps): JSX.Element {
  const [file, setFile] = useState<File | undefined>(undefined);

  const handleFileUpload: ChangeEventHandler<HTMLInputElement> = (e) => {
    if (e.target.files && e.target.files.length > 0) {
      setFile(e.target.files[0]);
    }

    if (onChange) onChange(e);
  };

  const handleSubmit: MouseEventHandler<HTMLButtonElement> = (e) => {
    if (file && onSubmit) onSubmit(file, e);
  };

  return (
    <div className={styles.container}>
      <input type="file" onChange={handleFileUpload} multiple={false} />
      <button type="submit" onClick={handleSubmit}>
        Upload
      </button>
    </div>
  );
}
