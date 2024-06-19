import Papa from "papaparse";
import { useFetchAllUserSolves } from "../../hooks/useFetch";
import Loading from "../Loading";

const exportData = (content: string, mimeType: string, extension: string) => {
  if (content) {
    const blob = new Blob([content], { type: mimeType });
    const url = URL.createObjectURL(blob);
    const link = document.createElement("a");
    link.download = `cubechron_solves.${extension}`;
    link.href = url;
    link.click();
    URL.revokeObjectURL(url);
  } else {
    console.error("No content");
  }
};

const FileExportMenu = () => {
  const { data: solves, isLoading } = useFetchAllUserSolves();
  if (isLoading || !solves) return <Loading />;
  return (
    <div className="flex flex-col gap-1">
      <button
        className="btn btn-xs"
        onClick={() =>
          exportData(JSON.stringify(solves), "application/json", "json")
        }
      >
        JSON
      </button>
      <button
        className="btn btn-xs"
        onClick={() => exportData(Papa.unparse(solves), "text/csv", "csv")}
      >
        CSV
      </button>
    </div>
  );
};

export default FileExportMenu;
