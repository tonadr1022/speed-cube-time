import { Solve, SolveCreatePayload } from "../types/types";
import { v4 as uuid } from "uuid";

const openDB = async () => {
  return new Promise<IDBDatabase>((resolve, reject) => {
    const req = indexedDB.open("timerDB", 1);

    req.onupgradeneeded = (event) => {
      const db = (event.target as IDBOpenDBRequest).result;
      if (!db.objectStoreNames.contains("solves")) {
        const store = db.createObjectStore("solves", {
          keyPath: "id",
        });
        store.createIndex("sessionId", "sessionId", { unique: false });
        store.createIndex("userId", "userId", { unique: false });
      }
    };
    req.onsuccess = () => {
      resolve(req.result);
    };

    req.onerror = (event) => {
      reject((event.target as IDBOpenDBRequest).error);
    };
  });
};

const withDB = async <T>(
  storeName: string,
  mode: IDBTransactionMode,
  callback: (store: IDBObjectStore) => IDBRequest,
): Promise<T> => {
  const db = await openDB();
  return new Promise<T>((resolve, reject) => {
    const transaction = db.transaction(storeName, mode);
    const store = transaction.objectStore(storeName);
    const request = callback(store);

    request.onsuccess = () => {
      resolve(request.result);
    };

    request.onerror = () => {
      reject(request.error);
    };
  });
};

export const fetchLocalSolves = async (): Promise<Solve[]> => {
  return withDB<Solve[]>("solves", "readonly", (store) => store.getAll());
};

export const createLocalSolve = async (solve: SolveCreatePayload) => {
  solve.id = uuid();
  return withDB<Solve>("solves", "readwrite", (store) => {
    const request = store.add(solve);
    return request;
  });
};
