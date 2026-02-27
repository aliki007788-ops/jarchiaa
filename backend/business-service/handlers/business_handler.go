package handlers

import (
    "encoding/json"
    "net/http"
    "business-service/models"
    "business-service/service"
    "github.com/gorilla/mux"
)

type BusinessHandler struct {
    planService *service.PlanService
}

func NewBusinessHandler(planService *service.PlanService) *BusinessHandler {
    return &BusinessHandler{
        planService: planService,
    }
}

// GET /api/business/plans
func (h *BusinessHandler) GetPlans(w http.ResponseWriter, r *http.Request) {
    plans, err := h.planService.GetActivePlans()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    json.NewEncoder(w).Encode(plans)
}

// POST /api/business/register
func (h *BusinessHandler) Register(w http.ResponseWriter, r *http.Request) {
    // دریافت user_id از توکن
    userID := r.Header.Get("X-User-ID")
    
    var req models.RegisterBusinessRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }
    
    business, err := h.planService.RegisterBusiness(userID, &req)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    json.NewEncoder(w).Encode(business)
}

// POST /api/business/{id}/purchase-plan
func (h *BusinessHandler) PurchasePlan(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    businessID := vars["id"]
    
    var req models.PurchasePlanRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }
    
    if err := h.planService.PurchasePlan(businessID, &req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{
        "message": "پلن با موفقیت خریداری شد",
    })
}

// GET /api/business/{id}/stats
func (h *BusinessHandler) GetStats(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    businessID := vars["id"]
    
    stats, err := h.planService.GetBusinessStats(businessID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    json.NewEncoder(w).Encode(stats)
}
