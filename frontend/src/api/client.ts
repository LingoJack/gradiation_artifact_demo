const BASE_URL = '/api/v1'

interface ApiResponse<T = any> {
  code: number
  message: string
  data: T
}

function getToken(): string | null {
  try {
    const raw = localStorage.getItem('user-storage')
    if (!raw) return null
    const parsed = JSON.parse(raw)
    return parsed?.state?.token || null
  } catch {
    return null
  }
}

async function request<T = any>(
  path: string,
  options: RequestInit = {},
): Promise<T> {
  const token = getToken()

  const headers: Record<string, string> = {
    'Content-Type': 'application/json',
    ...(options.headers as Record<string, string> || {}),
  }

  if (token) {
    headers['Authorization'] = `Bearer ${token}`
  }

  const res = await fetch(`${BASE_URL}${path}`, {
    ...options,
    headers,
  })

  const json: ApiResponse<T> = await res.json()

  if (json.code !== 0) {
    throw new Error(json.message || '请求失败')
  }

  return json.data
}

export function get<T = any>(path: string, params?: Record<string, any>): Promise<T> {
  let url = path
  if (params) {
    const qs = Object.entries(params)
      .filter(([, v]) => v !== undefined && v !== null && v !== '')
      .map(([k, v]) => `${encodeURIComponent(k)}=${encodeURIComponent(v)}`)
      .join('&')
    if (qs) url += `?${qs}`
  }
  return request<T>(url, { method: 'GET' })
}

export function post<T = any>(path: string, data?: any): Promise<T> {
  return request<T>(path, {
    method: 'POST',
    body: data !== undefined ? JSON.stringify(data) : undefined,
  })
}

export function put<T = any>(path: string, data?: any): Promise<T> {
  return request<T>(path, {
    method: 'PUT',
    body: data !== undefined ? JSON.stringify(data) : undefined,
  })
}

export function del<T = any>(path: string): Promise<T> {
  return request<T>(path, { method: 'DELETE' })
}
