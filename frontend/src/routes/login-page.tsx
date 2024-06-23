import { Link, useNavigate } from "react-router-dom";
import { toast } from "react-toastify";
import { useState } from "react";
import { maxCharMessage, minCharMessage } from "../util/validate-util";
import { ZodError, z } from "zod";
import { useAuth } from "../hooks/useContext";

const LoginUserData = z.object({
  username: z
    .string()
    .min(3, { message: minCharMessage(3) })
    .max(50, { message: maxCharMessage(50) }),
  password: z
    .string()
    .min(3, { message: minCharMessage(3) })
    .max(50, { message: maxCharMessage(50) }),
});

export default function LoginPage() {
  const auth = useAuth();
  const [data, setData] = useState({
    username: "",
    password: "",
  });

  const [errors, setErrors] = useState({
    username: "",
    password: "",
  });

  const navigate = useNavigate();
  const onLogin = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    try {
      LoginUserData.parse(data);
    } catch (e) {
      if (e instanceof ZodError) {
        const validationErrors = e.errors.reduce((acc: any, err) => {
          acc[err.path[0]] = err.message;
          return acc;
        }, {});
        setErrors((prevErrors) => ({
          ...prevErrors,
          ...validationErrors,
        }));
      }
    }

    try {
      await auth.login(data.username, data.password, () => navigate("/"));
    } catch (e) {
      if (e instanceof Error) {
        toast.error(e.message);
      }
    }
  };

  return (
    <>
      <div className="flex flex-col text-center justify-center items-center ">
        <div className="prose sm:mx-auto sm:w-full sm:max-w-sm">
          <img
            className="mx-auto h-16  w-auto"
            src="/pwa-512x512.png"
            alt="CubeChron Logo"
            width={30}
            height={30}
          />
          <h2 className="mt-2 text-center font-bold leading-9 tracking-tight  text-2xl">
            Sign In
          </h2>
        </div>
        <div className="sm:mx-auto sm:w-full sm:max-w-sm flex flex-col text-center">
          <form className="form-control space-y-4" onSubmit={onLogin}>
            <div>
              <label
                htmlFor="username"
                className="label text-sm text-start pb-0"
              >
                Username
              </label>
              <input
                id="username"
                name="username"
                autoComplete="username"
                value={data.username}
                onChange={(e) => setData({ ...data, username: e.target.value })}
                required
                className="input input-bordered input-primary w-full rounded-md "
              />
              {errors.username && (
                <p className="text-red-500 text-xs mt-1">{errors.username}</p>
              )}
            </div>

            <div>
              <label
                htmlFor="password"
                className="label text-sm text-start pb-0"
              >
                Password
              </label>
              <input
                id="password"
                name="password"
                type="password"
                autoComplete="current-password"
                required
                value={data.password}
                onChange={(e) => setData({ ...data, password: e.target.value })}
                className="input input-bordered input-primary w-full rounded-md"
              />
              {errors.password && (
                <p className="text-red-500 text-xs mt-1">{errors.password}</p>
              )}
            </div>
            <div>
              <button type="submit" className="btn btn-primary px-6">
                Sign in
              </button>
            </div>
            <div>
              No account?{" "}
              <Link className="link m-0" to="/register">
                Register
              </Link>
            </div>
          </form>
        </div>
      </div>
    </>
  );
}
