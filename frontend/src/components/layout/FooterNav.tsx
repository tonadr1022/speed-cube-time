import { FaGear, FaListUl } from "react-icons/fa6";
import { RxTimer } from "react-icons/rx";
import clsx from "clsx";
import { useLocation, useNavigate } from "react-router-dom";

const pathNames = [
  {
    icon: <RxTimer />,
    pathName: "/",
  },
  {
    icon: <FaListUl />,
    pathName: "/sessions",
  },
  {
    icon: <FaGear />,
    pathName: "/settings",
  },
];
type Props = {
  className?: string;
};

const FooterNav = ({ className }: Props) => {
  const loc = useLocation();
  const nav = useNavigate();

  return (
    <div className={clsx("flex flex-row w-full h-12  bg-base-300", className)}>
      {pathNames.map((pageData) => (
        <div key={pageData.pathName} className="w-full">
          <button
            onClick={() => nav(pageData.pathName)}
            className={clsx(
              "btn btn-sm h-full w-full p-0 m-0 rounded-none",
              pageData.pathName == loc.pathname && "btn-active",
            )}
          >
            {pageData.icon}
          </button>
        </div>
      ))}
    </div>
  );
};

export default FooterNav;
