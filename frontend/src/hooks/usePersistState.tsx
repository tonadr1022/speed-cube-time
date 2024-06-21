import { useState, useEffect } from "react";
import localforage from "localforage";

// FROM: https://jdudzik.medium.com/persistent-data-with-react-hooks-and-context-api-3f3f18ce947

// This hook receives two parameters:
// storageKey: This is the name of our storage that gets used when we retrieve/save our persistent data.
// initialState: This is our default value, but only if the store doesn't exist, otherwise it gets overwritten by the store.
const usePersistState = (
  storageKey: string,
  initialState: any,
  callback?: (newState: any) => void,
) => {
  const [state, setInternalState] = useState(initialState);

  // Create a replacement method that will set the state like normal, but that also saves the new state into the store.
  const setState = async (newState: any) => {
    await localforage.setItem(storageKey, newState);
    setInternalState(newState);
    if (callback) {
      callback(newState);
    }
  };

  // Only on our initial load, retrieve the data from the store and set the state to that data.
  useEffect(() => {
    const fetchData = async () => {
      try {
        const existing = await localforage.getItem(storageKey);
        if (existing !== null) {
          setInternalState(existing);
          if (callback) {
            callback(existing);
          }
        }
      } catch (error) {
        console.error("Failed to fetch data from localforage", error);
      }
    };

    fetchData();
  }, [storageKey, callback]);

  return [state, setState];
};

export default usePersistState;
