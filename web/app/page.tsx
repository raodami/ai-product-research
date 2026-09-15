export default function Home() {
  return (
    <main className="min-h-screen bg-gradient-to-br from-slate-900 via-purple-900 to-slate-900">
      <div className="max-w-6xl mx-auto px-4 py-20">
        <div className="text-center">
          <h1 className="text-5xl font-bold text-white mb-6">
            AI Product Research
          </h1>
          <p className="text-xl text-gray-300 mb-8">
            Discover opportunities. Analyze competitors. Find your edge.
          </p>
          <div className="flex gap-4 justify-center">
            <a href="/dashboard" className="bg-purple-600 hover:bg-purple-700 text-white px-8 py-3 rounded-xl font-medium transition-colors">
              Get Started
            </a>
          </div>
        </div>

        <div className="mt-20 grid md:grid-cols-3 gap-8">
          {[
            { icon: "🔍", title: "Scrape Sources", desc: "Toolify, Product Hunt, and more" },
            { icon: "📊", title: "Analyze Trends", desc: "AI-powered opportunity scoring" },
            { icon: "📈", title: "Export Reports", desc: "CSV, PDF, and detailed analysis" },
          ].map((feature, i) => (
            <div key={i} className="bg-slate-800/50 backdrop-blur rounded-2xl p-6 border border-slate-700">
              <div className="text-4xl mb-4">{feature.icon}</div>
              <h3 className="text-lg font-semibold text-white mb-2">{feature.title}</h3>
              <p className="text-gray-400">{feature.desc}</p>
            </div>
          ))}
        </div>
      </div>
    </main>
  );
}
