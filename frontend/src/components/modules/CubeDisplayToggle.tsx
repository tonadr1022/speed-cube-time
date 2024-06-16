type Props = {
  setIs3D: (is3D: boolean) => void;
  is3D: boolean;
};

const CubeDisplayToggle = ({ is3D, setIs3D }: Props) => {
  const handleUpdateSetting = (e: {
    currentTarget: { getAttribute: (arg0: string) => unknown };
  }) => {
    const value = e.currentTarget.getAttribute("value");
    setIs3D(value === "3D");
  };

  return (
    <div className="join text-center my-1">
      <input
        type="radio"
        name="options"
        data-title="3D"
        className="join-item btn btn-xs"
        value="3D"
        aria-label="3D"
        checked={is3D}
        onChange={(e) => handleUpdateSetting(e)}
      />
      <input
        type="radio"
        name="options"
        data-title="2D"
        className="join-item btn btn-xs"
        value="2D"
        aria-label="2D"
        checked={!is3D}
        onChange={(e) => handleUpdateSetting(e)}
      />
    </div>
  );
};

export default CubeDisplayToggle;
