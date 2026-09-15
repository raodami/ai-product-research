'use client';
import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import {
  ChartBarIcon,
  ArrowTrendingUpIcon,
  ArrowDownTrayIcon,
  PlusIcon,
  MagnifyingGlassIcon,
  FunnelIcon,
  CalendarIcon,
  SparklesIcon,
} from '@heroicons/react/24/outline';
import Link from 'next/link';
import { TrendLineChart, CategoryBarChart, OpportunityPieChart } from '@/components/charts';
import PWASetup from '@/components/pwa-setup';

interface Product {
  id: string;
  name: string;
  description: string;
  website: string;
  category: string;
  price: string;
  users: number;
}

interface Opportunity {
  keyword: string;
  market_size: string;
  competition: string;
  trend: string;
  opportunity_score: number;
  reasoning: string;
  suggestions: string[];
}

interface Trend {
  keyword: string;
  search_volume: number;
  growth_rate: number;
  opportunity: number;
}

export default function Dashboard() {
  const router = useRouter();
  const [products, setProducts] = useState<Product[]>([]);
  const [opportunities, setOpportunities] = useState<Opportunity[]>([]);
  const [trends, setTrends] = useState<Trend[]>([]);
  const [loading, setLoading] = useState(true);
  const [search, setSearch] = useState('');
  const [category, setCategory] = useState('');
  const [activeTab, setActiveTab] = useState<'products' | 'trends' | 'analytics'>('products');
  const [analysisKeyword, setAnalysisKeyword] = useState('');
  const [analyzing, setAnalyzing] = useState(false);
  const [schedule, setSchedule] = useState<string[]>([]);

  useEffect(() => {
    fetchData();
    fetchSchedule();
  }, []);

  const fetchData = async () => {
    try {
      const [productsRes, oppRes, trendsRes] = await Promise.all([
        fetch('/api/products?limit=50'),
        fetch('/api/trends'),
        fetch('/api/trends'),
      ]);
      
      if (productsRes.ok) {
        const data = await productsRes.json();
        setProducts(data);
      }
      
      if (oppRes.ok) {
        const data = await oppRes.json();
        setOpportunities(data);
      }

      if (trendsRes.ok) {
        const data = await trendsRes.json();
        setTrends(data);
      }
    } catch (error) {
      console.error('Failed to fetch data:', error);
    } finally {
      setLoading(false);
    }
  };

  const fetchSchedule = async () => {
    try {
      const res = await fetch('/api/scheduler/status');
      if (res.ok) {
        const data = await res.json();
        setSchedule(data.schedule);
      }
    } catch (error) {
      console.error('Failed to fetch schedule:', error);
    }
  };

  const handleSearch = async () => {
    setLoading(true);
    try {
      const params = new URLSearchParams();
      if (search) params.append('query', search);
      if (category) params.append('category', category);
      params.append('limit', '50');
      
      const res = await fetch(`/api/products?${params}`);
      if (res.ok) {
        const data = await res.json();
        setProducts(data);
      }
    } catch (error) {
      console.error('Search failed:', error);
    } finally {
      setLoading(false);
    }
  };

  const handleAnalyze = async () => {
    if (!analysisKeyword) return;
    
    setAnalyzing(true);
    try {
      const res = await fetch('/api/analyze/opportunity', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ keyword: analysisKeyword }),
      });
      
      if (res.ok) {
        const data = await res.json();
        setOpportunities([data, ...opportunities]);
        setAnalysisKeyword('');
      }
    } catch (error) {
      console.error('Analysis failed:', error);
    } finally {
      setAnalyzing(false);
    }
  };

  const handleExportCSV = async () => {
    window.open('/api/export/csv', '_blank');
  };

  const handleExportReport = async () => {
    window.open('/api/export/report', '_blank');
  };

  const handleRunScraper = async () => {
    try {
      await fetch('/api/scheduler/run', { method: 'POST' });
      alert('Scraper started! Check logs for results.');
    } catch (error) {
      console.error('Failed to run scraper:', error);
    }
  };

  const categories = ['All', 'AI Tools', 'SaaS', 'Dev Tools', 'Productivity', 'Marketing'];

  const categoryCounts = products.reduce((acc, p) => {
    acc[p.category] = (acc[p.category] || 0) + 1;
    return acc;
  }, {} as Record<string, number>);

  const categoryData = Object.entries(categoryCounts).map(([name, value]) => ({ name, value }));

  return (
    <div className="min-h-screen bg-gradient-to-br from-slate-900 via-purple-900 to-slate-900">
      <PWASetup />
      <div className="max-w-7xl mx-auto px-4 py-8">
        {/* Header */}
        <div className="flex justify-between items-center mb-8">
          <div>
            <h1 className="text-4xl font-bold text-white mb-2">AI Product Research</h1>
            <p className="text-purple-300">Discover opportunities and analyze competitors</p>
          </div>
          <Link
            href="/"
            className="px-6 py-3 bg-purple-600 hover:bg-purple-700 text-white rounded-lg transition-all"
          >
            Back Home
          </Link>
        </div>

        {/* Stats */}
        <div className="grid grid-cols-1 md:grid-cols-4 gap-6 mb-8">
          <div className="bg-slate-800/50 backdrop-blur rounded-xl p-6 border border-purple-500/20">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-purple-300 text-sm">Total Products</p>
                <p className="text-3xl font-bold text-white">{products.length}</p>
              </div>
              <ChartBarIcon className="w-10 h-10 text-purple-400" />
            </div>
          </div>
          
          <div className="bg-slate-800/50 backdrop-blur rounded-xl p-6 border border-purple-500/20">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-purple-300 text-sm">Trends Analyzed</p>
                <p className="text-3xl font-bold text-white">{opportunities.length}</p>
              </div>
              <ArrowTrendingUpIcon className="w-10 h-10 text-green-400" />
            </div>
          </div>
          
          <div className="bg-slate-800/50 backdrop-blur rounded-xl p-6 border border-purple-500/20">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-purple-300 text-sm">Scheduled Tasks</p>
                <p className="text-3xl font-bold text-white">{schedule.length}</p>
              </div>
              <CalendarIcon className="w-10 h-10 text-blue-400" />
            </div>
          </div>
          
          <div className="bg-slate-800/50 backdrop-blur rounded-xl p-6 border border-purple-500/20">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-purple-300 text-sm">Status</p>
                <p className="text-2xl font-bold text-green-400">Active</p>
              </div>
              <div className="w-4 h-4 bg-green-400 rounded-full animate-pulse" />
            </div>
          </div>
        </div>

        {/* Search and Filters */}
        <div className="bg-slate-800/50 backdrop-blur rounded-xl p-6 mb-6 border border-purple-500/20">
          <div className="flex flex-col md:flex-row gap-4">
            <div className="flex-1 relative">
              <MagnifyingGlassIcon className="absolute left-3 top-3 w-5 h-5 text-gray-400" />
              <input
                type="text"
                placeholder="Search products..."
                value={search}
                onChange={(e) => setSearch(e.target.value)}
                onKeyDown={(e) => e.key === 'Enter' && handleSearch()}
                className="w-full pl-10 pr-4 py-3 bg-slate-900/50 border border-purple-500/30 rounded-lg text-white placeholder-gray-400 focus:outline-none focus:border-purple-500"
              />
            </div>
            
            <div className="relative">
              <FunnelIcon className="absolute left-3 top-3 w-5 h-5 text-gray-400" />
              <select
                value={category}
                onChange={(e) => setCategory(e.target.value)}
                className="pl-10 pr-8 py-3 bg-slate-900/50 border border-purple-500/30 rounded-lg text-white focus:outline-none focus:border-purple-500 appearance-none"
              >
                <option value="">All Categories</option>
                {categories.filter(c => c !== 'All').map(c => (
                  <option key={c} value={c}>{c}</option>
                ))}
              </select>
            </div>
            
            <button
              onClick={handleSearch}
              className="px-6 py-3 bg-purple-600 hover:bg-purple-700 text-white rounded-lg transition-all font-medium"
            >
              Search
            </button>
          </div>
        </div>

        {/* Tabs */}
        <div className="flex gap-4 mb-6">
          <button
            onClick={() => setActiveTab('products')}
            className={`px-6 py-3 rounded-lg font-medium transition-all ${
              activeTab === 'products'
                ? 'bg-purple-600 text-white'
                : 'bg-slate-800/50 text-gray-400 hover:text-white'
            }`}
          >
            Products ({products.length})
          </button>
          <button
            onClick={() => setActiveTab('trends')}
            className={`px-6 py-3 rounded-lg font-medium transition-all ${
              activeTab === 'trends'
                ? 'bg-purple-600 text-white'
                : 'bg-slate-800/50 text-gray-400 hover:text-white'
            }`}
          >
            Trends & Opportunities ({opportunities.length})
          </button>
          <button
            onClick={() => setActiveTab('analytics')}
            className={`px-6 py-3 rounded-lg font-medium transition-all ${
              activeTab === 'analytics'
                ? 'bg-purple-600 text-white'
                : 'bg-slate-800/50 text-gray-400 hover:text-white'
            }`}
          >
            Analytics
          </button>
        </div>

        {/* Content */}
        {loading ? (
          <div className="text-center py-20">
            <div className="inline-block animate-spin rounded-full h-12 w-12 border-b-2 border-purple-400"></div>
            <p className="mt-4 text-gray-400">Loading...</p>
          </div>
        ) : activeTab === 'products' ? (
          <div className="bg-slate-800/50 backdrop-blur rounded-xl border border-purple-500/20 overflow-hidden">
            <div className="p-4 border-b border-purple-500/20 flex justify-between items-center">
              <h2 className="text-xl font-semibold text-white">Products</h2>
              <div className="flex gap-2">
                <button
                  onClick={handleExportCSV}
                  className="px-4 py-2 bg-green-600 hover:bg-green-700 text-white rounded-lg text-sm transition-all flex items-center gap-2"
                >
                  <ArrowDownTrayIcon className="w-4 h-4" />
                  Export CSV
                </button>
                <button
                  onClick={handleRunScraper}
                  className="px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded-lg text-sm transition-all flex items-center gap-2"
                >
                  <PlusIcon className="w-4 h-4" />
                  Run Scraper
                </button>
              </div>
            </div>
            
            <div className="overflow-x-auto">
              <table className="w-full">
                <thead className="bg-slate-900/50">
                  <tr>
                    <th className="px-6 py-3 text-left text-xs font-medium text-purple-300 uppercase">Product</th>
                    <th className="px-6 py-3 text-left text-xs font-medium text-purple-300 uppercase">Category</th>
                    <th className="px-6 py-3 text-left text-xs font-medium text-purple-300 uppercase">Price</th>
                    <th className="px-6 py-3 text-left text-xs font-medium text-purple-300 uppercase">Users</th>
                    <th className="px-6 py-3 text-left text-xs font-medium text-purple-300 uppercase">Website</th>
                  </tr>
                </thead>
                <tbody className="divide-y divide-purple-500/10">
                  {products.map((product) => (
                    <tr key={product.id} className="hover:bg-slate-700/30 transition-colors">
                      <td className="px-6 py-4">
                        <div className="font-medium text-white">{product.name}</div>
                        <div className="text-sm text-gray-400 truncate max-w-xs">{product.description}</div>
                      </td>
                      <td className="px-6 py-4">
                        <span className="px-3 py-1 bg-purple-500/20 text-purple-300 rounded-full text-sm">
                          {product.category}
                        </span>
                      </td>
                      <td className="px-6 py-4 text-gray-300">{product.price}</td>
                      <td className="px-6 py-4 text-gray-300">{product.users?.toLocaleString() ?? 'N/A'}</td>
                      <td className="px-6 py-4">
                        <a
                          href={product.website}
                          target="_blank"
                          rel="noopener noreferrer"
                          className="text-purple-400 hover:text-purple-300 underline text-sm"
                        >
                          Visit
                        </a>
                      </td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          </div>
        ) : activeTab === 'trends' ? (
          <div className="space-y-6">
            {/* Analysis Input */}
            <div className="bg-slate-800/50 backdrop-blur rounded-xl p-6 border border-purple-500/20">
              <h2 className="text-xl font-semibold text-white mb-4">Analyze Opportunity</h2>
              <div className="flex gap-4">
                <input
                  type="text"
                  placeholder="Enter keyword (e.g., AI writing tool)..."
                  value={analysisKeyword}
                  onChange={(e) => setAnalysisKeyword(e.target.value)}
                  onKeyDown={(e) => e.key === 'Enter' && handleAnalyze()}
                  className="flex-1 px-4 py-3 bg-slate-900/50 border border-purple-500/30 rounded-lg text-white placeholder-gray-400 focus:outline-none focus:border-purple-500"
                />
                <button
                  onClick={handleAnalyze}
                  disabled={analyzing || !analysisKeyword}
                  className="px-6 py-3 bg-purple-600 hover:bg-purple-700 disabled:bg-gray-600 disabled:cursor-not-allowed text-white rounded-lg transition-all font-medium flex items-center gap-2"
                >
                  <SparklesIcon className="w-5 h-5" />
                  {analyzing ? 'Analyzing...' : 'Analyze'}
                </button>
              </div>
            </div>

            {/* Schedule Info */}
            <div className="bg-slate-800/50 backdrop-blur rounded-xl p-6 border border-purple-500/20">
              <h3 className="text-lg font-semibold text-white mb-3 flex items-center gap-2">
                <CalendarIcon className="w-5 h-5 text-purple-400" />
                Scraper Schedule
              </h3>
              <ul className="space-y-2">
                {schedule.map((item, idx) => (
                  <li key={idx} className="flex items-center gap-2 text-gray-300">
                    <div className="w-2 h-2 bg-purple-400 rounded-full"></div>
                    {item}
                  </li>
                ))}
              </ul>
            </div>

            {/* Opportunities */}
            <div className="bg-slate-800/50 backdrop-blur rounded-xl border border-purple-500/20 overflow-hidden">
              <div className="p-4 border-b border-purple-500/20">
                <h2 className="text-xl font-semibold text-white">Opportunities</h2>
              </div>
              <div className="p-6 space-y-4">
                {opportunities.length === 0 ? (
                  <p className="text-gray-400 text-center py-8">No opportunities analyzed yet. Enter a keyword above.</p>
                ) : (
                  opportunities.map((opp, idx) => (
                    <div key={idx} className="bg-slate-900/50 rounded-lg p-4 border border-purple-500/10">
                      <div className="flex justify-between items-start mb-3">
                        <h3 className="text-lg font-semibold text-white">{opp.keyword}</h3>
                        <span className={`px-3 py-1 rounded-full text-sm font-medium ${
                          opp.opportunity_score >= 70 ? 'bg-green-500/20 text-green-400' :
                          opp.opportunity_score >= 50 ? 'bg-yellow-500/20 text-yellow-400' :
                          'bg-red-500/20 text-red-400'
                        }`}>
                          Score: {opp.opportunity_score}%
                        </span>
                      </div>
                      <div className="grid grid-cols-3 gap-4 mb-3">
                        <div>
                          <p className="text-purple-300 text-sm">Market Size</p>
                          <p className="text-white font-medium">{opp.market_size}</p>
                        </div>
                        <div>
                          <p className="text-purple-300 text-sm">Competition</p>
                          <p className="text-white font-medium">{opp.competition}</p>
                        </div>
                        <div>
                          <p className="text-purple-300 text-sm">Trend</p>
                          <p className="text-white font-medium">{opp.trend}</p>
                        </div>
                      </div>
                      <p className="text-gray-400 text-sm">{opp.reasoning}</p>
                      <div className="mt-3">
                        <p className="text-purple-300 text-sm mb-2">Suggestions:</p>
                        <ul className="list-disc list-inside text-gray-300 text-sm space-y-1">
                          {opp.suggestions.map((s, i) => <li key={i}>{s}</li>)}
                        </ul>
                      </div>
                    </div>
                  ))
                )}
              </div>
            </div>
          </div>
        ) : (
          <div className="space-y-6">
            {/* Charts Section */}
            <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
              {/* Trend Line Chart */}
              <div className="bg-slate-800/50 backdrop-blur rounded-xl p-6 border border-purple-500/20">
                <h3 className="text-lg font-semibold text-white mb-4">Search Trends</h3>
                {trends.length > 0 ? (
                  <TrendLineChart data={trends} />
                ) : (
                  <p className="text-gray-400 text-center py-8">No trend data available</p>
                )}
              </div>

              {/* Category Distribution */}
              <div className="bg-slate-800/50 backdrop-blur rounded-xl p-6 border border-purple-500/20">
                <h3 className="text-lg font-semibold text-white mb-4">Categories</h3>
                {categoryData.length > 0 ? (
                  <CategoryBarChart data={categoryData} />
                ) : (
                  <p className="text-gray-400 text-center py-8">No category data available</p>
                )}
              </div>
            </div>

            {/* Opportunity Distribution */}
            <div className="bg-slate-800/50 backdrop-blur rounded-xl p-6 border border-purple-500/20">
              <h3 className="text-lg font-semibold text-white mb-4">Opportunity Distribution</h3>
              {categoryData.length > 0 ? (
                <OpportunityPieChart data={categoryData} />
              ) : (
                <p className="text-gray-400 text-center py-8">No data available</p>
              )}
            </div>
          </div>
        )}
      </div>
    </div>
  );
}
