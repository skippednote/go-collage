const Home = () => {
  const [src, setSrc] = useState();
  const [err, setErr] = useState();

  const submit = async (e) => {
    e.preventDefault();
    let [uri, regex] = e.currentTarget.elements;
    uri = uri.value;
    regex = regex.value.replace(/\\/, "\\\\");

    try {
      const collageBlob = await fetch("/api/", {
        method: "POST",
        body: JSON.stringify({
          uri,
          regex,
        }),
      }).then((r) => r.blob());
      setSrc(URL.createObjectURL(collageBlob));
    } catch (e) {
      setErr(e.message);
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
