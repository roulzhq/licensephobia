export async function fileToBase64(file: File): Promise<string> {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.readAsDataURL(file);
    reader.onload = () => {
      if (reader.result != null && typeof reader.result === "string") {
        resolve(reader.result);
      } else {
        reject();
      }
    };
    reader.onerror = reject;
  });
}
