const BASE_URL = '/api/v1'

interface ApiResponse<T = unknown> {
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

  // 处理 HTTP 错误状态码
  if (!res.ok) {
    if (res.status === 401) {
      // 401 未授权，清除登录状态并跳转登录页
      localStorage.removeItem('user-storage')
      window.location.href = '/login'
      throw new Error('登录已过期，请重新登录')
    }
    throw new Error(`请求失败 (${res.status})，请稍后重试`)
  }

  // 安全解析 JSON，防止非 JSON 响应导致白屏
  let json: ApiResponse<T>
  try {
    json = await res.json()
  } catch {
    throw new Error('服务器响应异常，请稍后重试')
  }

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
