import { A } from "@solidjs/router";

export const App = (props: any) => (
  <>
    <nav>
      <A href="/">Home</A>
      <A href="/tasks">Tasks</A>
    </nav>
    <h1>BTCT</h1>
    {props.children}
  </>
);
