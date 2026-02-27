import { ref } from 'vue'
import axios from 'axios'

export function useVitrin() {
  const API_URL = 'https://jarchiaa.onrender.com/api'

  // دریافت آگهی‌های ویترین
  const getVitrinAds = async () => {
    try {
      const response = await axios.get(`${API_URL}/vitrin/ads`)
      return response.data
    } catch (error) {
      console.error('Error fetching vitrin ads:', error)
      // داده‌های نمونه برای تست
      return [
        {
          id: '1',
          title: 'پیتزا شیل',
          business_name: 'رستوران پیتزا شیل',
          business_id: 'b101',
          badge: 'ویژه',
          size: 'large',
          color: 'linear-gradient(135deg, #e84342, #ff7675)'
        },
        {
          id: '2',
          title: 'آرایشگاه مدرن',
          business_name: 'آرایشگاه مدرن',
          business_id: 'b102',
          badge: 'محبوب',
          size: 'vertical',
          color: 'linear-gradient(135deg, #6c5ce7, #a363d9)'
        }
      ]
    }
  }

  // دریافت جزئیات آگهی
  const getAdById = async (adId) => {
    try {
      const response = await axios.get(`${API_URL}/ads/${adId}`)
      return response.data
    } catch (error) {
      console.error('Error fetching ad:', error)
      // داده نمونه
      return {
        id: adId,
        title: 'پیتزا شیل',
        description: 'پیتزا خانواده با ۴۰٪ تخفیف ویژه',
        business_name: 'رستوران پیتزا شیل',
        business_id: 'b101',
        likes: 156,
        views: 1234,
        comments: [
          {
            id: 'c1',
            user_name: 'علی رضایی',
            text: 'عالی بود! حتماً امتحان کنید',
            time: '۲ ساعت پیش',
            likes: 5
          }
        ],
        color: 'linear-gradient(135deg, #e84342, #ff7675)'
      }
    }
  }

  // لایک کردن آگهی
  const likeAd = async (adId) => {
    const token = localStorage.getItem('token')
    await axios.post(`${API_URL}/ads/${adId}/like`, {}, {
      headers: { Authorization: `Bearer ${token}` }
    })
  }

  // افزودن نظر
  const addComment = async (adId, text) => {
    const token = localStorage.getItem('token')
    const response = await axios.post(`${API_URL}/ads/${adId}/comments`, 
      { text },
      { headers: { Authorization: `Bearer ${token}` } }
    )
    return response.data
  }

  return {
    getVitrinAds,
    getAdById,
    likeAd,
    addComment
  }
}