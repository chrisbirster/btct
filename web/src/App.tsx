import { createSignal } from "solid-js";

import { css } from "@linaria/core";

const button = css`
  padding: 10px 5px;
  border-radius: 10px;
  color: #bada55;
`;

const container = css`
  margin: 40px;
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: center;
`;

function App() {
  const [count, setCount] = createSignal(0);

  return (
    <div class={container}>
      <button class={button} onClick={() => setCount((count) => count - 1)}>
        -
      </button>
      <p>{count()}</p>
      <button class={button} onClick={() => setCount((count) => count + 1)}>
        +
      </button>
    </div>
  );
}

export default App;
