'use client';
import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import {
  ChartBarIcon,
  SearchIcon,
  ArrowTrendingUpIcon,
  DownloadIcon,
  PlusIcon,
} from '@heroicons/react/24/outline';

type Product = {
  id: string;
  name: string;
  category: string;
  price: string;
  users: string;
  score: number;
  website: string;
  description: string;
};

type Trend = {
  keyword: string;
  volume: number;
  growth: number;
  competition: string;
};

export default function Dashboard() {
  const router = useRouter();
  const [products, setProducts] = useState<Product[]>([]);
  const [trends, setTrends] = useState<Trend[]>([]);
  const [stats, setStats] = useState<any>(null);
  const [loading, setLoading] = useState(true);
  const [activeTab, setActiveTab] = useState('dashboard');

  useEffect(() => {
    Promise.all([
      fetch('/api/products').then(r => r.json()),
      fetch('/api/trends').then(r => r.json()),
      fetch('/api/analytics').then(r => r.json()),
    ]).then(([p, t, s]) => {
      setProducts(p);
      setTrends(t);
      setStats(s);
      setLoading(false);
    }).catch(() => setLoading(false));
  }, []);

  const handleScrape = async (source: string) => {
    await fetch(`/api/scraper/${source}`, { method: 'POST' });
    window.location.reload();
  };

  const handleExport = (format: string) => {
    window.open(`/api/export/${format}`, '_blank');
  };

  if (loading) {
    return (
      <div className="min-h-screen bg-slate-900 flex items-center justify-center">
        <div className="animate-spin h-8 w-8 border-2 border-purple-500 border-t-transparent rounded-full"></div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-slate-900">
      <header className="bg-slate-800/50 backdrop-blur border-b border-slate-700 sticky top-0 z-50">
        <div className="max-w-7xl mx-auto px-4 py-4 flex items-center justify-between">
          <h1 className="text-xl font-bold text-white">AI Product Research</h1>
          <div className="flex gap-2">
            <button onClick={() => handleExport('csv')} className="flex items-center gap-2 bg-slate-700 hover:bg-slate-600 px-3 py-2 rounded-lg text-sm">
              <DownloadIcon className="h-4 w-4" />
              CSV
            </button>
            <button onClick={() => handleExport('report')} className="flex items-center gap-2 bg-slate-700 hover:bg-slate-600 px-3 py-2 rounded-lg text-sm">
              <DownloadIcon className="h-4 w-4" />
              Report
            </button>
          </div>
        </div>
      </header>

      <main className="max-w-7xl mx-auto px-4 py-6">
        {/* Stats */}
        <div className="grid grid-cols-2 md:grid-cols-4 gap-4 mb-8">
          {[
            { label: 'Total Products', value: stats?.total_products || 0, icon: ChartBarIcon, color: 'text-blue-400' },
            { label: 'Avg Score', value: stats?.avg_score?.toFixed(1) || '0', icon: ArrowTrendingUpIcon, color: 'text-green-400' },
            { label: 'Trends', value: trends.length, icon: SearchIcon, color: 'text-purple-400' },
            { label: 'Categories', value: stats?.top_categories?.length || 0, icon: PlusIcon, color: 'text-yellow-400' },
          ].map((stat, i) => (
            <div key={i} className="bg-slate-800/50 rounded-xl p-4 border border-slate-700">
              <stat.icon className={`h-6 w-6 ${stat.color} mb-2`} />
              <p className="text-2xl font-bold text-white">{stat.value}</p>
              <p className="text-sm text-gray-400">{stat.label}</p>
            </div>
          ))}
        </div>

        {/* Tabs */}
        <div className="flex gap-2 mb-6">
          {['dashboard', 'products', 'trends', 'competitors'].map(tab => (
            <button
              key={tab}
              onClick={() => setActiveTab(tab)}
              className={`px-4 py-2 rounded-lg text-sm font-medium ${
                activeTab === tab ? 'bg-purple-600 text-white' : 'bg-slate-800 text-gray-400 hover:text-white'
              }`}
            >
              {tab.charAt(0).toUpperCase() + tab.slice(1)}
            </button>
          ))}
        </div>

        {/* Dashboard Tab */}
        {activeTab === 'dashboard' && (
          <div className="grid md:grid-cols-2 gap-6">
            <div className="bg-slate-800/50 rounded-xl p-6 border border-slate-700">
              <h2 className="text-lg font-semibold text-white mb-4">Top Products</h2>
              <div className="space-y-3">
                {products.slice(0, 5).map(p => (
                  <div key={p.id} className="flex items-center justify-between p-3 bg-slate-900/50 rounded-lg">
                    <div>
                      <p className="text-white font-medium">{p.name}</p>
                      <p className="text-sm text-gray-400">{p.category} • {p.price}</p>
                    </div>
                    <span className="text-purple-400 font-bold">{p.score.toFixed(1)}</span>
                  </div>
                ))}
              </div>
            </div>

            <div className="bg-slate-800/50 rounded-xl p-6 border border-slate-700">
              <h2 className="text-lg font-semibold text-white mb-4">Hot Trends</h2>
              <div className="space-y-3">
                {trends.slice(0, 5).map((t, i) => (
                  <div key={i} className="flex items-center justify-between p-3 bg-slate-900/50 rounded-lg">
                    <div>
                      <p className="text-white font-medium">{t.keyword}</p>
                      <p className="text-sm text-gray-400">{t.competition} competition</p>
                    </div>
                    <span className="text-green-400 font-bold">+{t.growth}%</span>
                  </div>
                ))}
              </div>
            </div>
          </div>
        )}

        {/* Products Tab */}
        {activeTab === 'products' && (
          <div className="bg-slate-800/50 rounded-xl border border-slate-700 overflow-hidden">
            <div className="p-4 border-b border-slate-700 flex justify-between items-center">
              <h2 className="text-lg font-semibold text-white">All Products</h2>
              <div className="flex gap-2">
                <button onClick={() => handleScrape('toolify')} className="bg-blue-600 hover:bg-blue-700 px-3 py-1 rounded text-sm">
                  Scrape Toolify
                </button>
                <button onClick={() => handleScrape('producthunt')} className="bg-orange-600 hover:bg-orange-700 px-3 py-1 rounded text-sm">
                  Scrape PH
                </button>
              </div>
            </div>
            <div className="divide-y divide-slate-700">
              {products.map(p => (
                <div key={p.id} className="p-4 hover:bg-slate-700/50 flex items-center justify-between">
                  <div>
                    <h3 className="text-white font-medium">{p.name}</h3>
                    <p className="text-sm text-gray-400">{p.description?.slice(0, 80)}...</p>
                    <p className="text-xs text-gray-500 mt-1">{p.category} • {p.price} • {p.users}</p>
                  </div>
                  <div className="text-right">
                    <span className="text-2xl font-bold text-purple-400">{p.score.toFixed(1)}</span>
                    <p className="text-xs text-gray-500">Score</p>
                  </div>
                </div>
              ))}
            </div>
          </div>
        )}

        {/* Trends Tab */}
        {activeTab === 'trends' && (
          <div className="bg-slate-800/50 rounded-xl border border-slate-700 overflow-hidden">
            <div className="p-4 border-b border-slate-700">
              <h2 className="text-lg font-semibold text-white">Trending Keywords</h2>
            </div>
            <div className="divide-y divide-slate-700">
              {trends.map((t, i) => (
                <div key={i} className="p-4 flex items-center justify-between hover:bg-slate-700/50">
                  <div>
                    <h3 className="text-white font-medium">{t.keyword}</h3>
                    <p className="text-sm text-gray-400">{t.competition} competition • {t.volume.toLocaleString()} searches</p>
                  </div>
                  <span className="text-green-400 font-bold text-lg">+{t.growth}%</span>
                </div>
              ))}
            </div>
          </div>
        )}
      </main>
    </div>
  );
}
