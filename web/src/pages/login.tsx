import { css } from "@linaria/core";

const loginButton = css`
  padding: 10px 20px;
  border-radius: 10px;
  color: white;
  background: #4285f4; /* Google-ish color */
`;

function Login() {
  return (
    <div>
      <h1>Please Log In</h1>
      <button
        class={loginButton}
        onClick={() => {
          window.location.href = "/auth/google";
        }}
      >
        Login with Google
      </button>
    </div>
  );
}

export default Login;
