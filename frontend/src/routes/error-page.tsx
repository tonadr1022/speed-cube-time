import { Link, useRouteError } from "react-router-dom";

export default function ErrorPage() {
  const error = useRouteError() as Error;

  return (
    <div className="flex flex-col justify-center text-center gap-8">
      <h2 className="mt-2 text-center font-bold leading-9 tracking-tight  text-2xl">
        Oops!
      </h2>
      <p>Sorry, an unexpected error has occurred.</p>
      <p>
        <i>Error: {error?.message}</i>
      </p>
      <button className="btn w-min self-center">
        <Link to="/">Home</Link>
      </button>
    </div>
  );
}
