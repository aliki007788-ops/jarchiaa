export function useAuth() {
  const login = async (phone) => {
    // کد قبلی
  }

  const verify = async (phone, code) => {
    try {
      // const response = await fetch('/api/auth/verify', ...)
      
      // برای تست:
      return {
        token: 'test-token',
        user: {
          id: '1',
          phone: phone,
          full_name: 'کاربر تست',
          role: 'user',
          isVerified: true
        }
      }
    } catch (error) {
      throw error
    }
  }

  const register = async (data) => {
    // کد قبلی
  }

  return { login, verify, register }
}
