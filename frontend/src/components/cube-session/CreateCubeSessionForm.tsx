import React, { useState } from "react";
import { CUBE_TYPE_OPTIONS } from "../../util/constants";
import { useQueryClient } from "@tanstack/react-query";
import { CubeSessionCreatePayload } from "../../types/types";
import { useCreateCubeSession } from "../../hooks/useFetch";

type Props = {
  onCompleted: () => void;
};

const initialFormState: CubeSessionCreatePayload = {
  name: "",
  cube_type: "333",
};

const CreateCubeSessionForm = ({ onCompleted }: Props) => {
  const [data, setData] = useState(initialFormState);
  const queryClient = useQueryClient();
  const createSessionMutation = useCreateCubeSession(queryClient);

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    createSessionMutation.mutate(data);
    onCompleted();
  };

  return (
    <div>
      <form
        className="flex flex-col justify-center gap-2"
        onSubmit={handleSubmit}
      >
        <h2 className="text-2xl font-bold pb-2">Create Cube Session</h2>
        <div className="w-full">
          <label htmlFor="name" className="label pb-0 font-medium">
            <span className="label-text text-base">Name</span>
          </label>
          <input
            className="input input-sm input-bordered w-full max-w-xs"
            type="text"
            value={data.name}
            autoFocus
            onChange={(e) => setData({ ...data, name: e.target.value })}
          />
        </div>
        <div>
          <label htmlFor="cubeType" className="label pb-0 font-medium">
            <span className="text-base label-text">Cube Type</span>
          </label>
          <select
            className="select select-sm select-bordered w-full max-w-xs"
            name="cubeType"
            value={data.cube_type}
            onChange={(e) => setData({ ...data, cube_type: e.target.value })}
          >
            {Object.entries(CUBE_TYPE_OPTIONS).map(([key, value]) => (
              <option key={key} value={key}>
                {value}
              </option>
            ))}
          </select>
        </div>
        <button
          className="btn btn-success text-center self-center"
          type="submit"
        >
          Create
        </button>
      </form>
    </div>
  );
};

export default CreateCubeSessionForm;
