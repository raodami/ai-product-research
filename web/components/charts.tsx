'use client';

import {
  LineChart,
  Line,
  BarChart,
  Bar,
  PieChart,
  Pie,
  Cell,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  Legend,
  ResponsiveContainer,
} from 'recharts';

interface Trend {
  keyword: string;
  search_volume: number;
  growth_rate: number;
  opportunity: number;
}

interface CategoryData {
  name: string;
  value: number;
}

export function TrendLineChart({ data }: { data: Trend[] }) {
  const chartData = data.map((d) => ({
    name: d.keyword.slice(0, 10),
    volume: d.search_volume,
    growth: d.growth_rate,
    opportunity: d.opportunity,
  }));

  return (
    <ResponsiveContainer width="100%" height={300}>
      <LineChart data={chartData}>
        <CartesianGrid strokeDasharray="3 3" stroke="rgba(99, 102, 241, 0.2)" />
        <XAxis dataKey="name" stroke="#94a3b8" fontSize={12} />
        <YAxis stroke="#94a3b8" fontSize={12} />
        <Tooltip 
          contentStyle={{ 
            backgroundColor: 'rgba(15, 23, 42, 0.95)', 
            border: '1px solid rgba(99, 102, 241, 0.3)',
            borderRadius: '8px'
          }}
        />
        <Legend />
        <Line type="monotone" dataKey="volume" stroke="#533afd" strokeWidth={2} name="Search Volume" />
        <Line type="monotone" dataKey="growth" stroke="#10b981" strokeWidth={2} name="Growth Rate" />
        <Line type="monotone" dataKey="opportunity" stroke="#f59e0b" strokeWidth={2} name="Opportunity" />
      </LineChart>
    </ResponsiveContainer>
  );
}

export function CategoryBarChart({ data }: { data: CategoryData[] }) {
  return (
    <ResponsiveContainer width="100%" height={300}>
      <BarChart data={data}>
        <CartesianGrid strokeDasharray="3 3" stroke="rgba(99, 102, 241, 0.2)" />
        <XAxis dataKey="name" stroke="#94a3b8" fontSize={12} />
        <YAxis stroke="#94a3b8" fontSize={12} />
        <Tooltip 
          contentStyle={{ 
            backgroundColor: 'rgba(15, 23, 42, 0.95)', 
            border: '1px solid rgba(99, 102, 241, 0.3)',
            borderRadius: '8px'
          }}
        />
        <Bar dataKey="value" fill="#533afd" radius={[4, 4, 0, 0]} />
      </BarChart>
    </ResponsiveContainer>
  );
}

const COLORS = ['#533afd', '#10b981', '#f59e0b', '#ef4444', '#8b5cf6', '#06b6d4'];

export function OpportunityPieChart({ data }: { data: CategoryData[] }) {
  return (
    <ResponsiveContainer width="100%" height={300}>
      <PieChart>
        <Pie
          data={data}
          cx="50%"
          cy="50%"
          labelLine={false}
          label={({ name, percent }) => `${name}: ${((percent ?? 0) * 100).toFixed(0)}%`}
          outerRadius={100}
          fill="#8884d8"
          dataKey="value"
        >
          {data.map((entry, index) => (
            <Cell key={`cell-${index}`} fill={COLORS[index % COLORS.length]} />
          ))}
        </Pie>
        <Tooltip 
          contentStyle={{ 
            backgroundColor: 'rgba(15, 23, 42, 0.95)', 
            border: '1px solid rgba(99, 102, 241, 0.3)',
            borderRadius: '8px'
          }}
        />
      </PieChart>
    </ResponsiveContainer>
  );
}
