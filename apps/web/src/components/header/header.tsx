import Link from "next/link";
import styles from "./header.module.scss";

export default function Header(): JSX.Element {
  return (
    <header className={styles.header}>
      <Link href="/">
        <h1 className={styles.title}>Licensephobia</h1>
      </Link>
      <nav className={styles.nav}>
        <ul>
          <li>
            <Link href="/licenses">Licenses</Link>
          </li>
          <li>
            <Link href="/about">About</Link>
          </li>
        </ul>
      </nav>
    </header>
  );
}
