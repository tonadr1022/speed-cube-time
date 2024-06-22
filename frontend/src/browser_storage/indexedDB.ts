import {
  CubeSession,
  CubeSessionCreatePayload,
  Settings,
  Solve,
  SolveCreatePayload,
} from "../types/types";
import { v4 as uuid } from "uuid";

const openDB = async () => {
  return new Promise<IDBDatabase>((resolve, reject) => {
    try {
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

        if (!db.objectStoreNames.contains("settings")) {
          db.createObjectStore("settings", {
            keyPath: "id",
          });
        }
        if (!db.objectStoreNames.contains("cubeSessions")) {
          db.createObjectStore("cubeSessions", {
            keyPath: "id",
          });
        }
      };

      req.onsuccess = () => {
        resolve(req.result);
      };

      req.onerror = (event) => {
        reject((event.target as IDBOpenDBRequest).error);
      };
    } catch (e) {
      console.log("Error: ", e);
      throw e;
    }
  });
};

const withDB = async <T>(
  storeName: string,
  mode: IDBTransactionMode,
  callback: (store: IDBObjectStore) => IDBRequest,
): Promise<T> => {
  const db = await openDB();
  return new Promise<T>((resolve, reject) => {
    try {
      const transaction = db.transaction(storeName, mode);
      const store = transaction.objectStore(storeName);
      const request = callback(store);

      request.onsuccess = () => {
        resolve(request.result);
      };

      request.onerror = () => {
        reject(request.error);
      };
    } catch (e) {
      console.log("error: ", e);
    }
  });
};

export const fetchLocalSolves = async (): Promise<Solve[]> => {
  return withDB<Solve[]>("solves", "readonly", (store) => store.getAll());
};

export const fetchLocalCubeSessions = async (): Promise<CubeSession[]> => {
  return withDB<CubeSession[]>("cubeSessions", "readonly", (store) =>
    store.getAll(),
  );
};

export const fetchLocalSettings = async (): Promise<Settings | undefined> => {
  return withDB<Settings | undefined>("settings", "readonly", (store) => {
    return store.getAll();
  })
    .then((settingsArray) => {
      if (settingsArray && settingsArray.length > 0) {
        return settingsArray[0]; // Return the first (and only) settings object
      } else {
        return undefined; // No settings found
      }
    })
    .catch((error) => {
      console.error("Error fetching user settings:", error);
      return undefined; // Handle error and return undefined
    });
};

export const updateLocalSettings = async (
  settings: Partial<Settings>,
): Promise<void> => {
  return withDB<void>("settings", "readwrite", (store) => {
    return new Promise<void>((resolve, reject) => {
      const getRequest = store.getAll();
      getRequest.onsuccess = (event) => {
        const settingsArray = event.target.result;
        let updateRequest: IDBRequest;
        if (settingsArray && settingsArray.length > 0) {
          const existingSettings = settingsArray[0];
          for (const key in settings) {
            if (Object.prototype.hasOwnProperty.call(settings, key)) {
              existingSettings[key] = settings[key];
            }
          }
          updateRequest = store.put(existingSettings);
        } else {
          console.log("add new settings", settings);
          updateRequest = store.add(settings as Settings);
        }

        updateRequest.onsuccess = () => {
          resolve();
        };

        updateRequest.onerror = () => {
          reject("Failed to update user settings");
        };
      };

      getRequest.onerror = () => {
        reject("Error fetching user settings");
      };
    });
  });
};

export const updateLocalSolve = async (id: string, solve: Partial<Solve>) => {
  return withDB<void>("solves", "readwrite", (store) => {
    return new Promise<void>((resolve, reject) => {
      const getRequest = store.get(id);
      getRequest.onsuccess = (event) => {
        const existing = event.target.result;
        if (existing) {
          for (const key in existing) {
            if (Object.prototype.hasOwnProperty.call(solve, key)) {
              existing[key] = solve[key];
            }
          }
          const updateRequest = store.put(existing);
          updateRequest.onsuccess = () => {
            resolve();
          };
          updateRequest.onerror = () => {
            reject("Failed to update solve");
          };
        } else {
          reject("solve doesn't exist");
        }
      };
      getRequest.onerror = () => {
        reject("Error fetching existing solve");
      };
    });
  });
};

export const createLocalSettings = async (settings: Settings) => {
  return withDB<Settings>("settings", "readwrite", (store) => {
    return store.add(settings);
  });
};

export const createLocalCubeSession = async (
  cubeSession: CubeSessionCreatePayload,
) => {
  const date = new Date();
  const session: CubeSession = {
    ...cubeSession,
    id: uuid(),
    created_at: date,
    updated_at: date,
  };
  return withDB<CubeSession>("cubeSessions", "readwrite", (store) => {
    return store.add(session);
  });
};

export const createLocalSolve = async (solve: SolveCreatePayload) => {
  solve.id = uuid();
  return withDB<Solve>("solves", "readwrite", (store) => {
    const request = store.add(solve);
    return request;
  });
};

export const deleteLocalSolve = async (id: string) => {
  return withDB<Solve>("solves", "readwrite", (store) => {
    return store.delete(id);
  });
};
