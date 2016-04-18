Vim�UnDo� � C,�L9)���H��l=�����4M$�   G                                  W��    _�                     ;       ����                                                                                                                                                                                                                                                                                                                                                             V��     �   :   <   G      .	if err != nil && err == gorm.RecordNotFound {5�_�                    C   '    ����                                                                                                                                                                                                                                                                                                                                                             V��     �   B   D   G      D		goa.Error(ctx, "updating portfolio", goa.KV{"error", err.Error()})5�_�                    C   '    ����                                                                                                                                                                                                                                                                                                                                                             V��     �   B   D   G      C		goa.Error(ctx, "updating portfolio", oa.KV{"error", err.Error()})5�_�                    C   '    ����                                                                                                                                                                                                                                                                                                                                                             V��     �   B   D   G      B		goa.Error(ctx, "updating portfolio", a.KV{"error", err.Error()})5�_�                    C   '    ����                                                                                                                                                                                                                                                                                                                                                             V��     �   B   D   G      A		goa.Error(ctx, "updating portfolio", .KV{"error", err.Error()})5�_�                    C   '    ����                                                                                                                                                                                                                                                                                                                                                             V��     �   B   D   G      @		goa.Error(ctx, "updating portfolio", KV{"error", err.Error()})5�_�                    C   '    ����                                                                                                                                                                                                                                                                                                                                                             V��     �   B   D   G      ?		goa.Error(ctx, "updating portfolio", V{"error", err.Error()})5�_�      	              C   '    ����                                                                                                                                                                                                                                                                                                                                                             V��     �   B   D   G      >		goa.Error(ctx, "updating portfolio", {"error", err.Error()})5�_�      
           	   C   '    ����                                                                                                                                                                                                                                                                                                                                                             V��     �   B   D   G      =		goa.Error(ctx, "updating portfolio", "error", err.Error()})5�_�   	              
   C   '    ����                                                                                                                                                                                                                                                                                                                                                             V��     �   B   D   G      <		goa.Error(ctx, "updating portfolio", error", err.Error()})5�_�   
                 C   '    ����                                                                                                                                                                                                                                                                                                                                                             V��     �   B   D   G      ;		goa.Error(ctx, "updating portfolio", rror", err.Error()})5�_�                    C   '    ����                                                                                                                                                                                                                                                                                                                                                             V��     �   B   D   G      :		goa.Error(ctx, "updating portfolio", ror", err.Error()})5�_�                    C   '    ����                                                                                                                                                                                                                                                                                                                                                             V��     �   B   D   G      9		goa.Error(ctx, "updating portfolio", or", err.Error()})5�_�                    C   '    ����                                                                                                                                                                                                                                                                                                                                                             V��     �   B   D   G      8		goa.Error(ctx, "updating portfolio", r", err.Error()})5�_�                    C   '    ����                                                                                                                                                                                                                                                                                                                                                             V��     �   B   D   G      7		goa.Error(ctx, "updating portfolio", ", err.Error()})5�_�                    C   '    ����                                                                                                                                                                                                                                                                                                                                                             V��     �   B   D   G      6		goa.Error(ctx, "updating portfolio", , err.Error()})5�_�                    C   '    ����                                                                                                                                                                                                                                                                                                                                                             V��     �   B   D   G      5		goa.Error(ctx, "updating portfolio",  err.Error()})5�_�                    C   2    ����                                                                                                                                                                                                                                                                                                                                                             V��     �   B   D   G      4		goa.Error(ctx, "updating portfolio", err.Error()})5�_�                    C   2    ����                                                                                                                                                                                                                                                                                                                                                             V��    �               G   package main       import (   	"xor/cmd/api/app"   	"xor/cmd/api/models"       	"github.com/goadesign/goa"   	"github.com/jinzhu/gorm"   )       8// PortfolioController implements theportfolio resource.   !type PortfolioController struct {   	*goa.Controller   	e           *Env   	storagelist map[string]Storage   $	storage     models.PortfolioStorage   }       9// NewPortfolioController creates a portfolio controller.   ofunc NewPortfolioController(service goa.Service, storagelist map[string]Storage, e *Env) *PortfolioController {   	return &PortfolioController{   <		Controller:  service.NewController("PortfolioController"),   		e:           e,   		storagelist: storagelist,   I		storage:     storagelist["PORTFOLIOSTORAGE"].(models.PortfolioStorage),   	}   }       !// Create runs the create action.   Mfunc (c *PortfolioController) Create(ctx *app.CreatePortfolioContext) error {   ?	mem := models.PortfolioFromCreatePortfolioPayload(ctx.Payload)   -	rmem, err := c.storage.Add(ctx.Context, mem)   	if err != nil {   6		return goa.Response(ctx).Send(ctx, 500, err.Error())   	}       T	ctx.ResponseData.Header().Set("Location", app.PortfolioHref(ctx.MemberID, rmem.ID))   	return ctx.Created()   }       !// Delete runs the delete action.   Mfunc (c *PortfolioController) Delete(ctx *app.DeletePortfolioContext) error {   6	err := c.storage.Delete(ctx.Context, ctx.PortfolioID)   	if err != nil {   		return ctx.NotFound()   	}   	return ctx.NoContent()   }       // List runs the list action.   Ifunc (c *PortfolioController) List(ctx *app.ListPortfolioContext) error {   :	res := c.storage.ListPortfolio(ctx.Context, ctx.MemberID)   	return ctx.OK(res)   }       // Show runs the show action.   Ifunc (c *PortfolioController) Show(ctx *app.ShowPortfolioContext) error {   O	res, err := c.storage.OnePortfolio(ctx.Context, ctx.PortfolioID, ctx.MemberID)   1	if err != nil && err == gorm.ErrRecordNotFound {   		return ctx.NotFound()   	}   	return ctx.OK(res)   #} // Update runs the update action.   Mfunc (c *PortfolioController) Update(ctx *app.UpdatePortfolioContext) error {   ]	err := c.storage.UpdateFromUpdatePortfolioPayload(ctx.Context, ctx.Payload, ctx.PortfolioID)   	if err != nil {   3		goa.Error(ctx, "updating portfolio", err.Error())   		return ctx.Err()   	}   	return ctx.NoContent()   }5�_�                    "       ����                                                                                                                                                                                                                                                                                                                                                             V���     �   !   #   G      6		return goa.Response(ctx).Send(ctx, 500, err.Error())5�_�                    C       ����                                                                                                                                                                                                                                                                                                                                                             V���     �   B   D   G      3		goa.Error(ctx, "updating portfolio", err.Error())5�_�                    C       ����                                                                                                                                                                                                                                                                                                                                                             V���    �               G   package main       import (   	"xor/cmd/api/app"   	"xor/cmd/api/models"       	"github.com/goadesign/goa"   	"github.com/jinzhu/gorm"   )       8// PortfolioController implements theportfolio resource.   !type PortfolioController struct {   	*goa.Controller   	e           *Env   	storagelist map[string]Storage   $	storage     models.PortfolioStorage   }       9// NewPortfolioController creates a portfolio controller.   ofunc NewPortfolioController(service goa.Service, storagelist map[string]Storage, e *Env) *PortfolioController {   	return &PortfolioController{   <		Controller:  service.NewController("PortfolioController"),   		e:           e,   		storagelist: storagelist,   I		storage:     storagelist["PORTFOLIOSTORAGE"].(models.PortfolioStorage),   	}   }       !// Create runs the create action.   Mfunc (c *PortfolioController) Create(ctx *app.CreatePortfolioContext) error {   ?	mem := models.PortfolioFromCreatePortfolioPayload(ctx.Payload)   -	rmem, err := c.storage.Add(ctx.Context, mem)   	if err != nil {   =		return goa.ContextResponse(ctx).Send(ctx, 500, err.Error())   	}       T	ctx.ResponseData.Header().Set("Location", app.PortfolioHref(ctx.MemberID, rmem.ID))   	return ctx.Created()   }       !// Delete runs the delete action.   Mfunc (c *PortfolioController) Delete(ctx *app.DeletePortfolioContext) error {   6	err := c.storage.Delete(ctx.Context, ctx.PortfolioID)   	if err != nil {   		return ctx.NotFound()   	}   	return ctx.NoContent()   }       // List runs the list action.   Ifunc (c *PortfolioController) List(ctx *app.ListPortfolioContext) error {   :	res := c.storage.ListPortfolio(ctx.Context, ctx.MemberID)   	return ctx.OK(res)   }       // Show runs the show action.   Ifunc (c *PortfolioController) Show(ctx *app.ShowPortfolioContext) error {   O	res, err := c.storage.OnePortfolio(ctx.Context, ctx.PortfolioID, ctx.MemberID)   1	if err != nil && err == gorm.ErrRecordNotFound {   		return ctx.NotFound()   	}   	return ctx.OK(res)   #} // Update runs the update action.   Mfunc (c *PortfolioController) Update(ctx *app.UpdatePortfolioContext) error {   ]	err := c.storage.UpdateFromUpdatePortfolioPayload(ctx.Context, ctx.Payload, ctx.PortfolioID)   	if err != nil {   6		goa.LogError(ctx, "updating portfolio", err.Error())   		return ctx.Err()   	}   	return ctx.NoContent()   }5�_�                    !       ����                                                                                                                                                                                                                                                                                                                                                             W��     �       #   G      	if err != nil {5�_�                    "        ����                                                                                                                                                                                                                                                                                                                                                             W��     �   "   $   H    �   "   #   H    5�_�                    $       ����                                                                                                                                                                                                                                                                                                                                                             W��     �   #   $          =		return goa.ContextResponse(ctx).Send(ctx, 500, err.Error())5�_�                    "        ����                                                                                                                                                                                                                                                                                                                                                             W��     �   !   "           5�_�                     "       ����                                                                                                                                                                                                                                                                                                                                                             W��    �               G   package main       import (   	"xor/cmd/api/app"   	"xor/cmd/api/models"       	"github.com/goadesign/goa"   	"github.com/jinzhu/gorm"   )       8// PortfolioController implements theportfolio resource.   !type PortfolioController struct {   	*goa.Controller   	e           *Env   	storagelist map[string]Storage   $	storage     models.PortfolioStorage   }       9// NewPortfolioController creates a portfolio controller.   ofunc NewPortfolioController(service goa.Service, storagelist map[string]Storage, e *Env) *PortfolioController {   	return &PortfolioController{   <		Controller:  service.NewController("PortfolioController"),   		e:           e,   		storagelist: storagelist,   I		storage:     storagelist["PORTFOLIOSTORAGE"].(models.PortfolioStorage),   	}   }       !// Create runs the create action.   Mfunc (c *PortfolioController) Create(ctx *app.CreatePortfolioContext) error {   ?	mem := models.PortfolioFromCreatePortfolioPayload(ctx.Payload)   -	rmem, err := c.storage.Add(ctx.Context, mem)   	if err != nil {   0		return ctx.Service.Send(ctx, 500, err.Error())   	}       T	ctx.ResponseData.Header().Set("Location", app.PortfolioHref(ctx.MemberID, rmem.ID))   	return ctx.Created()   }       !// Delete runs the delete action.   Mfunc (c *PortfolioController) Delete(ctx *app.DeletePortfolioContext) error {   6	err := c.storage.Delete(ctx.Context, ctx.PortfolioID)   	if err != nil {   		return ctx.NotFound()   	}   	return ctx.NoContent()   }       // List runs the list action.   Ifunc (c *PortfolioController) List(ctx *app.ListPortfolioContext) error {   :	res := c.storage.ListPortfolio(ctx.Context, ctx.MemberID)   	return ctx.OK(res)   }       // Show runs the show action.   Ifunc (c *PortfolioController) Show(ctx *app.ShowPortfolioContext) error {   O	res, err := c.storage.OnePortfolio(ctx.Context, ctx.PortfolioID, ctx.MemberID)   1	if err != nil && err == gorm.ErrRecordNotFound {   		return ctx.NotFound()   	}   	return ctx.OK(res)   #} // Update runs the update action.   Mfunc (c *PortfolioController) Update(ctx *app.UpdatePortfolioContext) error {   ]	err := c.storage.UpdateFromUpdatePortfolioPayload(ctx.Context, ctx.Payload, ctx.PortfolioID)   	if err != nil {   6		goa.LogError(ctx, "updating portfolio", err.Error())   		return ctx.Err()   	}   	return ctx.NoContent()   }5��