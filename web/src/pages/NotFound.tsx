import { useParams } from "@solidjs/router";

export default function NotFound() {
  const params = useParams();
  return (
    <div>
      <h1>BTCT</h1>
      <p>Error retrieving route</p>
      <p>{JSON.stringify(params)}</p>
    </div>
  );
}
