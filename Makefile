#


# calc for next_version according last_version
#	$(1) nxt_version - result of upgrade version
#	$(2) kind 	- major/minor/patch/build
#	$(3) stage 	- alpha/beta/rc
define f_upgrade_version
	$(eval _kind=$(if $(strip $(2)),$(strip $(2)),build))

	$(eval last_version=$(shell git describe --tags --abbrev=0))
	$(eval version_core=$(word 1,$(subst -, ,$(last_version:v%=%))))
	$(eval majorX=$(word 1,$(subst ., ,$(version_core))))
	$(eval minorY=$(word 2,$(subst ., ,$(version_core))))
	$(eval patchZ=$(word 3,$(subst ., ,$(version_core))))

	$(eval version_suffix=$(word 2,$(subst -, ,$(last_version))))
	$(eval prerelease=$(word 1,$(subst ., ,$(version_suffix))))

	$(eval stageT=$(word 1,$(subst ., ,$(version_suffix))))

	$(eval buildT=$(word 2,$(subst ., ,$(version_suffix))))
	$(eval buildT=$(if $(strip $(buildT)),$(buildT),0))
	$(eval buildT=$(if $(and $(filter-out major minor patch,_kind),$(strip $(3)),$(filter $(3),$(stageT))),$(shell expr "$(buildT)" + 1),1))

	$(eval majorX=$(if $(filter major,$(2)),$(shell expr "$(majorX)" + 1),$(majorX)))
	$(eval minorY=$(if $(filter minor,$(2)),$(shell expr "$(minorY)" + 1),$(minorY)))
	$(eval patchZ=$(if $(filter patch,$(2)),$(shell expr "$(patchZ)" + 1),$(patchZ)))
	$(eval suffixT=$(if $(filter alpha beta rc,$(3)),-$(3).$(buildT),))

	$(1)=v$(majorX).$(minorY).$(patchZ)$(suffixT)
endef


.pony: publish
publish:
	$(eval kind=$(shell read -p "Enter the kind of publish(major/minor/patch/build):" kind; echo $${kind}))
	$(eval stage=$(shell read -p "Enter the stage of release(alpha/beta/rc/release):" stage; if [ stage == "release" ]; then stage=""; fi; echo $${stage}))
	$(eval $(call f_upgrade_version,nxt_version,$(kind),$(stage)))
	$(eval confirm=$(shell read -p "Are you sure to publish { $(nxt_version) } <$(kind)>@<$(stage)> (y or n)?" value; echo $${value}))
	@echo $(nxt_version) > VERSION
	git tag "$(nxt_version)"
	git push origin "$(nxt_version)"
	go list -m github.com/shawnwy/go-utils@$(nxt_version)
