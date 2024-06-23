import { toast } from "react-toastify";
import { useState } from "react";
import { Link, useNavigate } from "react-router-dom";

import { ZodError, z } from "zod";
import { minCharMessage, maxCharMessage } from "../util/validate-util";
import { useAuth } from "../hooks/useContext";

const passwordMatchErrorMsg = "Passwords do not match";
const passwordZString = z
  .string()
  .min(3, { message: minCharMessage(3) })
  .max(100, { message: maxCharMessage(100) });

const RegisterUserData = z.object({
  username: z
    .string()
    .min(3, { message: minCharMessage(3) })
    .max(50, { message: maxCharMessage(50) }),
  password: passwordZString,
  passwordConfirm: passwordZString,
});

export default function RegisterPage() {
  const auth = useAuth();
  const navigate = useNavigate();
  const [data, setData] = useState({
    username: "",
    password: "",
    passwordConfirm: "",
  });

  const [errors, setErrors] = useState({
    passwordConfirm: "",
    username: "",
    password: "",
  });

  const registerUser = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    if (errors.passwordConfirm) {
      toast.error(passwordMatchErrorMsg);
      return;
    }
    try {
      RegisterUserData.parse(data);
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
        const errors = e.errors.map((err) => err.message).join(", ");
        toast.error(`${errors}`);
        return;
      }
    }
    try {
      await auth.register(data.username, data.password, () =>
        navigate("/login"),
      );
    } catch (error) {
      if (error instanceof Error) {
        toast.error(error.message);
      }
    }
  };

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setData((prevData) => {
      const newData = { ...prevData, [name]: value };
      const isPasswordMatchError =
        newData.password.length > 0 &&
        newData.passwordConfirm.length > 0 &&
        newData.password !== newData.passwordConfirm;
      setErrors((prevErrors) => {
        return {
          ...prevErrors,
          passwordConfirm: isPasswordMatchError ? passwordMatchErrorMsg : "",
        };
      });
      return newData;
    });
  };

  return (
    <>
      <div className="flex flex-col text-center justify-center items-center ">
        <div className="prose sm:mx-auto sm:w-full sm:max-w-sm">
          <img
            className="mx-auto h-16 w-auto m-0"
            src="/pwa-512x512.png"
            alt="CubeChron Logo"
            width={30}
            height={30}
          />
          <h2 className="mt-2 text-center font-bold leading-9 tracking-tight  text-2xl">
            Register
          </h2>
        </div>
        <form className="space-y-4" onSubmit={registerUser}>
          <div>
            <label
              htmlFor="username"
              className="label block text-sm text-start pb-0"
            >
              Username
            </label>
            <input
              id="username"
              name="username"
              type="text"
              value={data.username}
              onChange={handleChange}
              required
              className="input input-bordered input-primary block w-full rounded-md "
            />
            {errors.username && (
              <p className="text-red-500 text-xs mt-1">{errors.username}</p>
            )}
          </div>
          <div>
            <label htmlFor="password" className="label text-sm text-start pb-0">
              Password
            </label>
            <input
              id="password"
              name="password"
              type="password"
              autoComplete="current-password"
              required
              value={data.password}
              onChange={handleChange}
              className="input input-bordered input-primary w-full rounded-md"
            />
            {errors.password && (
              <p className="text-red-500 text-xs mt-1">{errors.password}</p>
            )}
          </div>
          <div>
            <label
              htmlFor="passwordConfirm"
              className="label text-sm text-start pb-0"
            >
              Confirm Password
            </label>
            <input
              id="passwordConfirm"
              name="passwordConfirm"
              type="password"
              autoComplete="current-password"
              required
              value={data.passwordConfirm}
              onChange={handleChange}
              className="input input-bordered input-primary w-full rounded-md"
            />
            {errors.passwordConfirm && (
              <p className="text-red-500 text-xs mt-1">
                {errors.passwordConfirm}
              </p>
            )}
          </div>
          <div></div>
          <div>
            <button type="submit" className="btn btn-primary px-6">
              Register
            </button>
          </div>
          <div>
            Have an Account?{" "}
            <Link className="link" to="/login">
              Login
            </Link>
          </div>
        </form>
      </div>
    </>
  );
}
