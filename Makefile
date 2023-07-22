#


# calc for next_version according last_version
#	$(1) nxt_version - result of upgrade version $(shell git describe --tags --abbrev=0)
#	$(2) kind 	- major/minor/patch/build
#	$(3) stage 	- alpha/beta/rc
define f_upgrade_version
	$(eval _kind=$(if $(strip $(2)),$(strip $(2)),build))
	$(info _kind)

	$(eval last_version=v5.0.1-alpha.1)
	$(eval version_core=$(word 1,$(subst -, ,$(last_version:v%=%))))
	$(eval majorX=$(word 1,$(subst ., ,$(version_core))))
	$(eval minorY=$(word 2,$(subst ., ,$(version_core))))
	$(eval patchZ=$(word 3,$(subst ., ,$(version_core))))

	$(eval version_suffix=$(word 2,$(subst -, ,$(last_version))))
	$(eval prerelease=$(word 1,$(subst ., ,$(version_suffix))))

	$(eval stageT=$(word 1,$(subst ., ,$(version_suffix))))
	$(info "stageT - $(stageT)")

	$(eval buildT=$(word 2,$(subst ., ,$(version_suffix))))
	$(eval buildT=$(if $(strip $(buildT)),$(buildT),0))
	$(eval buildT=$(if $(and $(filter-out major minor patch,_kind),$(strip $(3)),$(filter $(3),$(stageT))),$(shell expr "$(buildT)" + 1),1))
	$(info "test->$(buildT)")

	$(eval majorX=$(if $(filter major,$(2)),$(shell expr "$(majorX)" + 1),$(majorX)))
	$(eval minorY=$(if $(filter minor,$(2)),$(shell expr "$(minorY)" + 1),$(minorY)))
	$(eval patchZ=$(if $(filter patch,$(2)),$(shell expr "$(patchZ)" + 1),$(patchZ)))
	$(eval suffixT=$(if $(filter alpha beta rc,$(3)),-$(3).$(buildT),))
	$(info "suffix- $(suffixT)")

	$(1)=v$(majorX).$(minorY).$(patchZ)$(suffixT)
	$(info "version - v$(majorX).$(minorY).$(patchZ)$(suffixT)")
endef

define f_x
	$(eval $(1)=$(if $(1),,build))
	$(info $(1))
endef

.pony: ver
ver:
	$(eval kind=$(shell read -p "Enter the kind of publish(major/minor/patch/build):" kind; echo $${kind}))
	$(eval stage=$(shell read -p "Enter the stage of release(alpha/beta/rc/release):" stage; if [ stage == "release" ]; then stage=""; fi; echo $${stage}))
	$(eval $(call f_upgrade_version,nxt_version,$(kind),$(stage)))
	$(eval confirm=$(shell read -p "Are you sure to publish { $(next_version) } <$(kind)>@<$(stage)> (y or n)?" value; echo $${value}))
	echo $(nxt_version) > VERSION
	git tag "$(nxt_version)"
	git push origin "$(nxt_version)"
	GOPROXY=proxy.golang.org go list -m github.com/shawnwy/go-utils@$(nxt_version)


test:
	$(eval version_suffix=alpha.1)
	$(eval s=alpha)
	$(eval stage=$(word 1,$(subst ., ,$(version_suffix))))
	$(eval mode=build)
	$(info "s=$(s) stage=$(stage)")
	@echo $(strip $(s))
	@echo $(filter $(s),$(stage))
	@echo $(filter major minor patch,$(mode))
	@echo $(if $(and $(strip $(s)),$(filter $(s),$(stage))),$(shell expr "$(buildT)" + 1),1)


test2:
	$(eval confirm=y)
	$(info $(strip $(confirm)))
ifneq ($(confirm),y)
	$(error aborted action)
else
	echo "no end"
endif
	echo "end"

ver2:
	$(info $(filter-out major minor patch,))
	$(if $(filter-out major minor patch,build),true,false)

MODE ?= patch
.pony: publish
publish:
	$(call f_upgrade_version,minor) > VERSION
	$(eval nxt_version=`cat VERSION`)
	git tag "$(nxt_version)"
	git push origin "$(nxt_version)"
	GOPROXY=proxy.golang.org go list -m github.com/fasionchan/goutils@$version
