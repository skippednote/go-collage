import { useState } from "react";
const Home = () => {
  const [src, setSrc] = useState();
  const [err, setErr] = useState();

  const submit = async (e) => {
    e.preventDefault();
    let [uri, regex, width, gray] = e.currentTarget.elements;
    uri = uri.value;
    regex = regex.value.replace(/\\/, "\\\\");
    width = width.value;
    gray = gray.value;

    try {
      const response = await fetch("/api/", {
        method: "POST",
        body: JSON.stringify({
          uri,
          regex,
          width,
          gray,
        }),
      });
      const collageBlob = await response.blob();
      setSrc(URL.createObjectURL(collageBlob));
    } catch (e) {
      setErr("Failed to load the image");
      console.log(e);
    }
  };
  return (
    <main>
      <h1>Home</h1>
      {err && <h4>{err}</h4>}
      <form onSubmit={submit}>
        <input placeholder="url" />
        <input placeholder="image regex" />
        <input placeholder="width" />
        <input placeholder="gray?" type="checkbox" />
        <button type="submit">Submit</button>
      </form>
      {src && (
        <>
          <hr />
          <img src={src} />
        </>
      )}
    </main>
  );
};
export default Home;
