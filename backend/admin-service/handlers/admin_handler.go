package handlers

import (
    "admin-service/models"
    "admin-service/service"
    "encoding/json"
    "net/http"
    "strconv"
    "github.com/gorilla/mux"
)

type AdminHandler struct {
    moduleService *service.ModuleService
}

func NewAdminHandler(moduleService *service.ModuleService) *AdminHandler {
    return &AdminHandler{
        moduleService: moduleService,
    }
}

// GET /api/admin/stats
func (h *AdminHandler) GetStats(w http.ResponseWriter, r *http.Request) {
    stats, err := h.moduleService.GetSystemStats()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(stats)
}

// GET /api/admin/modules
func (h *AdminHandler) GetModules(w http.ResponseWriter, r *http.Request) {
    modules, err := h.moduleService.GetAllModules()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(modules)
}

// PUT /api/admin/modules/{key}/toggle
func (h *AdminHandler) ToggleModule(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    key := vars["key"]
    
    var req struct {
        Active bool `json:"active"`
    }
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }
    
    if err := h.moduleService.ToggleModule(key, req.Active); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{
        "message": "وضعیت ماژول با موفقیت تغییر کرد",
    })
}

// POST /api/admin/plans
func (h *AdminHandler) CreatePlan(w http.ResponseWriter, r *http.Request) {
    var plan models.PlanConfig
    if err := json.NewDecoder(r.Body).Decode(&plan); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }
    
    if err := h.moduleService.CreatePlan(&plan); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(plan)
}

// PUT /api/admin/plans/{id}
func (h *AdminHandler) UpdatePlan(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, _ := strconv.Atoi(vars["id"])
    
    var plan models.PlanConfig
    if err := json.NewDecoder(r.Body).Decode(&plan); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }
    
    if err := h.moduleService.UpdatePlan(id, &plan); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    json.NewEncoder(w).Encode(plan)
}

// PUT /api/admin/businesses/{id}/verify
func (h *AdminHandler) VerifyBusiness(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    businessID := vars["id"]
    
    var req struct {
        Verified bool `json:"verified"`
    }
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }
    
    if err := h.moduleService.VerifyBusiness(businessID, req.Verified); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{
        "message": "وضعیت کسب‌وکار به‌روزرسانی شد",
    })
}

// GET /api/admin/config/{key}
func (h *AdminHandler) GetConfig(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    key := vars["key"]
    
    value, err := h.moduleService.GetConfig(key)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    
    json.NewEncoder(w).Encode(map[string]string{
        "key":   key,
        "value": value,
    })
}
