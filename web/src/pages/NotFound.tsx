import { A, useParams } from "@solidjs/router";

export default function NotFound() {
  const params = useParams();
  return (
    <>
      <nav>
        <A href="/">Home</A>
        <A href="/tasks">Tasks</A>
      </nav>
      <h1>BTCT</h1>
      <p>Error retrieving route</p>
      <p>{JSON.stringify(params)}</p>
    </>
  );
}
