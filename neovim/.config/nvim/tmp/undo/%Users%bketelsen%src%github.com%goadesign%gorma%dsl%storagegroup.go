Vim�UnDo� �Q�UYtd��vˋ���YC�A�	8��B �-q   )                                  V�G�    _�                            ����                                                                                                                                                                                                                                                                                                                                                             V�G�     �   
      )      )	if !dslengine.TopLevelDefinition(true) {5�_�                       $    ����                                                                                                                                                                                                                                                                                                                                                             V�G�     �   
      )      +	if !dslengine.IsTopLevelDefinition(true) {5�_�                        $    ����                                                                                                                                                                                                                                                                                                                                                             V�G�    �               )   package dsl       import (   %	"github.com/goadesign/goa/dslengine"   	"github.com/goadesign/gorma"   )       2// StorageGroup implements the top level Gorma DSL   8// There should be one StorageGroup per Goa application.   Kfunc StorageGroup(name string, dsli func()) *gorma.StorageGroupDefinition {   '	if !dslengine.IsTopLevelDefinition() {   		return nil   	}   	if name == "" {   =		dslengine.ReportError("Storage Group name cannot be empty")   	}       	if gorma.GormaDesign != nil {   %		if gorma.GormaDesign.Name == name {   <			dslengine.ReportError("Only one StorageGroup is allowed")   		}   	}   	gorma.GormaDesign.Name = name   W	gorma.GormaDesign.RelationalStores = make(map[string]*gorma.RelationalStoreDefinition)   '	gorma.GormaDesign.DefinitionDSL = dsli   	return gorma.GormaDesign   }       /// Description sets the definition description.   c// Description can be called inside StorageGroup, RelationalStore, RelationalModel, RelationalField   func Description(d string) {   0	if a, ok := storageGroupDefinition(false); ok {   		a.Description = d   :	} else if v, ok := relationalStoreDefinition(false); ok {   		v.Description = d   :	} else if r, ok := relationalModelDefinition(false); ok {   		r.Description = d   :	} else if f, ok := relationalFieldDefinition(false); ok {   		f.Description = d   	}   }5��