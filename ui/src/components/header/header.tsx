function Header() {
  return (
    <>
      <header style={{ backgroundColor: "var(--foreground)" }}>
        <h1 style={{ color: "var(--accent)" }} className="p-6 font-bold text-4xl">
          Weather App
        </h1>
      </header>
    </>
  );
}

export { Header };
