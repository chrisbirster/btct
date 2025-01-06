import { css } from "@linaria/core";

const container = css`
  margin: 40px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  gap: 1rem;
  font-size: 2rem;
`;

export default function Task() {
  return (
    <div class={container}>
      <h1>Task</h1>
    </div>
  );
}
