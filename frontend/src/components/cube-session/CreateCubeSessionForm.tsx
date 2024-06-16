import React, { useState } from "react";
import { CUBE_TYPE_OPTIONS } from "../../util/constants";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { createCubeSession } from "../../api/cube-session-api";
import { CubeSession, CubeSessionCreatePayload } from "../../types/types";

type Props = {
  onCompleted: () => void;
};

const initialFormState: CubeSessionCreatePayload = {
  name: "",
  cubeType: "333",
};

const CreateCubeSessionForm = ({ onCompleted }: Props) => {
  const [data, setData] = useState(initialFormState);
  const queryClient = useQueryClient();
  const createSessionMutation = useMutation({
    mutationFn: (session: CubeSessionCreatePayload) =>
      createCubeSession(session, false),
    onMutate: async (newSession: CubeSessionCreatePayload) => {
      await queryClient.cancelQueries({ queryKey: ["cubeSessions"] });
      const prevSessions = queryClient.getQueryData(["cubeSessions"]);
      queryClient.setQueryData(["cubeSessions"], (old: CubeSession[]) => [
        ...old,
        newSession,
      ]);
      return { prevSessions };
    },
    onError: (_, __, context) => {
      queryClient.setQueryData(["cubeSessions"], context?.prevSessions);
    },
    onSettled: () => {
      queryClient.invalidateQueries({ queryKey: ["cubeSessions"] });
      onCompleted();
    },
  });

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    createSessionMutation.mutate(data);

    // createCubeSession(data);
    // updateSetting(setting!, {
    //   cubeType: data.cubeType!,
    //   id: setting!.setting.id,
    // });
  };

  return (
    <div>
      <form onSubmit={handleSubmit}>
        <h2 className="text-2xl font-bold pb-2">Create Cube Session</h2>
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
        <label htmlFor="cubeType" className="label pb-0 font-medium">
          <span className="text-base label-text">Cube Type</span>
        </label>
        <select
          className="select select-sm select-bordered w-full max-w-xs"
          name="cubeType"
          value={data.cubeType}
          onChange={(e) => setData({ ...data, cubeType: e.target.value })}
        >
          {Object.entries(CUBE_TYPE_OPTIONS).map(([key, value]) => (
            <option key={key} value={key}>
              {value}
            </option>
          ))}
        </select>
        <button className="btn btn-primary text-center" type="submit">
          Create
        </button>
      </form>
    </div>
  );
};

export default CreateCubeSessionForm;
