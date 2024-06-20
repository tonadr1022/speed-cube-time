import { useSettings } from "../../hooks/useContext";
import { MODULE_OPTIONS } from "../../util/constants";
import { handleDropdownOptionClick } from "../../util/handleDropdownClick";

type Props = {
  moduleNumber: number;
};

const ModuleSelect = ({ moduleNumber }: Props) => {
  const { modules, setModules } = useSettings();
  const handleModuleSelect = (module: string) => {
    const newModules = modules.map((m, i) => {
      if (i === moduleNumber) {
        return module;
      }
      return m;
    });
    setModules(newModules);
    handleDropdownOptionClick();
  };

  return (
    <div className="dropdown">
      <button
        tabIndex={0}
        className="m-1 btn btn-xs bg-primary text-primary-content hover:bg-primary"
      >
        {modules[moduleNumber]}
      </button>
      <ul
        tabIndex={0}
        className="p-0 menu dropdown-content border-base-300 border-2 bg-base-100 rounded-box w-40 max-h-64 overflow-y-auto block"
      >
        {Object.entries(MODULE_OPTIONS).map(([key, value]) => (
          <li
            value={key}
            key={key}
            onClick={() => handleModuleSelect(key)}
            className="hover:bg-base-300"
          >
            <button>{value}</button>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default ModuleSelect;
