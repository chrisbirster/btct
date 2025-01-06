import { A } from "@solidjs/router";

export default function Home() {
  return (
    <>
      <nav>
        <A href="/">Home</A>
        <A href="/tasks">Tasks</A>
      </nav>
      <h1>BTCT</h1>
      <p>Home</p>
    </>
  );
}
