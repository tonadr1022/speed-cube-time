import { handleDropdownOptionClick } from "../../util/handleDropdownClick";
import { FaChevronDown } from "react-icons/fa6";
import { useSettings } from "../../hooks/useContext";

const moduleCounts = [1, 2, 3, 4];

const ModuleCountSelect = () => {
  const { moduleCount, setModuleCount } = useSettings();
  const handleSettingUpdate = (count: number) => {
    handleDropdownOptionClick();
    setModuleCount(count);
  };
  return (
    <>
      <div className="dropdown dropdown-end ">
        <div tabIndex={0} className="m-1 btn btn-sm bg-base-300">
          {moduleCount} <FaChevronDown />
        </div>
        <ul
          tabIndex={0}
          className="p-2 menu  dropdown-content bg-base-100 rounded-box w-40 max-h-64 overflow-y-auto block"
        >
          {moduleCounts.map((i) => (
            <li value={i} key={i} onClick={() => handleSettingUpdate(i)}>
              <a className="hover:bg-base-300">{i}</a>
            </li>
          ))}
        </ul>
      </div>
    </>
  );
};

export default ModuleCountSelect;
