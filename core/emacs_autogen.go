// Automatically generated with pimacs.extract
// DO NOT MODIFY!
// Generated from GNU Emacs commit: 4b6f2a7028b91128934a19f83572f24106782225, branch: emacs-29

package core

const emacsConstant_charset_arg_max_autogen = 17
const emacsConstant_coding_arg_max_autogen = 13

func (ec *execContext) boolVector_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("bool-vector") // Source file: alloc.c
}

func (ec *execContext) cons_autogen(car, cdr lispObject) (lispObject, error) {
	return ec.stub("cons") // Source file: alloc.c
}

func (ec *execContext) garbageCollect_autogen() (lispObject, error) {
	return ec.stub("garbage-collect") // Source file: alloc.c
}

func (ec *execContext) garbageCollectMaybe_autogen(factor lispObject) (lispObject, error) {
	return ec.stub("garbage-collect-maybe") // Source file: alloc.c
}

func (ec *execContext) list_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("list") // Source file: alloc.c
}

func (ec *execContext) makeBoolVector_autogen(length, init lispObject) (lispObject, error) {
	return ec.stub("make-bool-vector") // Source file: alloc.c
}

func (ec *execContext) makeByteCode_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("make-byte-code") // Source file: alloc.c
}

func (ec *execContext) makeClosure_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("make-closure") // Source file: alloc.c
}

func (ec *execContext) makeFinalizer_autogen(function lispObject) (lispObject, error) {
	return ec.stub("make-finalizer") // Source file: alloc.c
}

func (ec *execContext) makeList_autogen(length, init lispObject) (lispObject, error) {
	return ec.stub("make-list") // Source file: alloc.c
}

func (ec *execContext) makeMarker_autogen() (lispObject, error) {
	return ec.stub("make-marker") // Source file: alloc.c
}

func (ec *execContext) makeRecord_autogen(type_, slots, init lispObject) (lispObject, error) {
	return ec.stub("make-record") // Source file: alloc.c
}

func (ec *execContext) makeString_autogen(length, init, multibyte lispObject) (lispObject, error) {
	return ec.stub("make-string") // Source file: alloc.c
}

func (ec *execContext) makeSymbol_autogen(name lispObject) (lispObject, error) {
	return ec.stub("make-symbol") // Source file: alloc.c
}

func (ec *execContext) makeVector_autogen(length, init lispObject) (lispObject, error) {
	return ec.stub("make-vector") // Source file: alloc.c
}

func (ec *execContext) mallocInfo_autogen() (lispObject, error) {
	return ec.stub("malloc-info") // Source file: alloc.c
}

func (ec *execContext) mallocTrim_autogen(leave_padding lispObject) (lispObject, error) {
	return ec.stub("malloc-trim") // Source file: alloc.c
}

func (ec *execContext) memoryInfo_autogen() (lispObject, error) {
	return ec.stub("memory-info") // Source file: alloc.c
}

func (ec *execContext) memoryUseCounts_autogen() (lispObject, error) {
	return ec.stub("memory-use-counts") // Source file: alloc.c
}

func (ec *execContext) purecopy_autogen(obj lispObject) (lispObject, error) {
	return ec.stub("purecopy") // Source file: alloc.c
}

func (ec *execContext) record_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("record") // Source file: alloc.c
}

func (ec *execContext) suspiciousObject_autogen(obj lispObject) (lispObject, error) {
	return ec.stub("suspicious-object") // Source file: alloc.c
}

func (ec *execContext) vector_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("vector") // Source file: alloc.c
}

func (ec *execContext) debugTimerCheck_autogen() (lispObject, error) {
	return ec.stub("debug-timer-check") // Source file: atimer.c
}

func (ec *execContext) barfIfBufferReadOnly_autogen(position lispObject) (lispObject, error) {
	return ec.stub("barf-if-buffer-read-only") // Source file: buffer.c
}

func (ec *execContext) bufferBaseBuffer_autogen(buffer lispObject) (lispObject, error) {
	return ec.stub("buffer-base-buffer") // Source file: buffer.c
}

func (ec *execContext) bufferCharsModifiedTick_autogen(buffer lispObject) (lispObject, error) {
	return ec.stub("buffer-chars-modified-tick") // Source file: buffer.c
}

func (ec *execContext) bufferEnableUndo_autogen(buffer lispObject) (lispObject, error) {
	return ec.stub("buffer-enable-undo") // Source file: buffer.c
}

func (ec *execContext) bufferFileName_autogen(buffer lispObject) (lispObject, error) {
	return ec.stub("buffer-file-name") // Source file: buffer.c
}

func (ec *execContext) bufferList_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("buffer-list") // Source file: buffer.c
}

func (ec *execContext) bufferLiveP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("buffer-live-p") // Source file: buffer.c
}

func (ec *execContext) bufferLocalValue_autogen(variable, buffer lispObject) (lispObject, error) {
	return ec.stub("buffer-local-value") // Source file: buffer.c
}

func (ec *execContext) bufferLocalVariables_autogen(buffer lispObject) (lispObject, error) {
	return ec.stub("buffer-local-variables") // Source file: buffer.c
}

func (ec *execContext) bufferModifiedP_autogen(buffer lispObject) (lispObject, error) {
	return ec.stub("buffer-modified-p") // Source file: buffer.c
}

func (ec *execContext) bufferModifiedTick_autogen(buffer lispObject) (lispObject, error) {
	return ec.stub("buffer-modified-tick") // Source file: buffer.c
}

func (ec *execContext) bufferName_autogen(buffer lispObject) (lispObject, error) {
	return ec.stub("buffer-name") // Source file: buffer.c
}

func (ec *execContext) bufferSwapText_autogen(buffer lispObject) (lispObject, error) {
	return ec.stub("buffer-swap-text") // Source file: buffer.c
}

func (ec *execContext) buryBufferInternal_autogen(buffer lispObject) (lispObject, error) {
	return ec.stub("bury-buffer-internal") // Source file: buffer.c
}

func (ec *execContext) currentBuffer_autogen() (lispObject, error) {
	return ec.stub("current-buffer") // Source file: buffer.c
}

func (ec *execContext) deleteAllOverlays_autogen(buffer lispObject) (lispObject, error) {
	return ec.stub("delete-all-overlays") // Source file: buffer.c
}

func (ec *execContext) deleteOverlay_autogen(overlay lispObject) (lispObject, error) {
	return ec.stub("delete-overlay") // Source file: buffer.c
}

func (ec *execContext) eraseBuffer_autogen() (lispObject, error) {
	return ec.stub("erase-buffer") // Source file: buffer.c
}

func (ec *execContext) forceModeLineUpdate_autogen(all lispObject) (lispObject, error) {
	return ec.stub("force-mode-line-update") // Source file: buffer.c
}

func (ec *execContext) generateNewBufferName_autogen(name, ignore lispObject) (lispObject, error) {
	return ec.stub("generate-new-buffer-name") // Source file: buffer.c
}

func (ec *execContext) getBuffer_autogen(buffer_or_name lispObject) (lispObject, error) {
	return ec.stub("get-buffer") // Source file: buffer.c
}

func (ec *execContext) getBufferCreate_autogen(buffer_or_name, inhibit_buffer_hooks lispObject) (lispObject, error) {
	return ec.stub("get-buffer-create") // Source file: buffer.c
}

func (ec *execContext) getFileBuffer_autogen(filename lispObject) (lispObject, error) {
	return ec.stub("get-file-buffer") // Source file: buffer.c
}

func (ec *execContext) internalSetBufferModifiedTick_autogen(tick, buffer lispObject) (lispObject, error) {
	return ec.stub("internal--set-buffer-modified-tick") // Source file: buffer.c
}

func (ec *execContext) killAllLocalVariables_autogen(kill_permanent lispObject) (lispObject, error) {
	return ec.stub("kill-all-local-variables") // Source file: buffer.c
}

func (ec *execContext) killBuffer_autogen(buffer_or_name lispObject) (lispObject, error) {
	return ec.stub("kill-buffer") // Source file: buffer.c
}

func (ec *execContext) makeIndirectBuffer_autogen(base_buffer, name, clone, inhibit_buffer_hooks lispObject) (lispObject, error) {
	return ec.stub("make-indirect-buffer") // Source file: buffer.c
}

func (ec *execContext) makeOverlay_autogen(beg, end, buffer, front_advance, rear_advance lispObject) (lispObject, error) {
	return ec.stub("make-overlay") // Source file: buffer.c
}

func (ec *execContext) moveOverlay_autogen(overlay, beg, end, buffer lispObject) (lispObject, error) {
	return ec.stub("move-overlay") // Source file: buffer.c
}

func (ec *execContext) nextOverlayChange_autogen(pos lispObject) (lispObject, error) {
	return ec.stub("next-overlay-change") // Source file: buffer.c
}

func (ec *execContext) otherBuffer_autogen(buffer, visible_ok, frame lispObject) (lispObject, error) {
	return ec.stub("other-buffer") // Source file: buffer.c
}

func (ec *execContext) overlayBuffer_autogen(overlay lispObject) (lispObject, error) {
	return ec.stub("overlay-buffer") // Source file: buffer.c
}

func (ec *execContext) overlayEnd_autogen(overlay lispObject) (lispObject, error) {
	return ec.stub("overlay-end") // Source file: buffer.c
}

func (ec *execContext) overlayGet_autogen(overlay, prop lispObject) (lispObject, error) {
	return ec.stub("overlay-get") // Source file: buffer.c
}

func (ec *execContext) overlayLists_autogen() (lispObject, error) {
	return ec.stub("overlay-lists") // Source file: buffer.c
}

func (ec *execContext) overlayProperties_autogen(overlay lispObject) (lispObject, error) {
	return ec.stub("overlay-properties") // Source file: buffer.c
}

func (ec *execContext) overlayPut_autogen(overlay, prop, value lispObject) (lispObject, error) {
	return ec.stub("overlay-put") // Source file: buffer.c
}

func (ec *execContext) overlayRecenter_autogen(pos lispObject) (lispObject, error) {
	return ec.stub("overlay-recenter") // Source file: buffer.c
}

func (ec *execContext) overlayStart_autogen(overlay lispObject) (lispObject, error) {
	return ec.stub("overlay-start") // Source file: buffer.c
}

func (ec *execContext) overlayTree_autogen(buffer lispObject) (lispObject, error) {
	return ec.stub("overlay-tree") // Source file: buffer.c
}

func (ec *execContext) overlayp_autogen(object lispObject) (lispObject, error) {
	return ec.stub("overlayp") // Source file: buffer.c
}

func (ec *execContext) overlaysAt_autogen(pos, sorted lispObject) (lispObject, error) {
	return ec.stub("overlays-at") // Source file: buffer.c
}

func (ec *execContext) overlaysIn_autogen(beg, end lispObject) (lispObject, error) {
	return ec.stub("overlays-in") // Source file: buffer.c
}

func (ec *execContext) previousOverlayChange_autogen(pos lispObject) (lispObject, error) {
	return ec.stub("previous-overlay-change") // Source file: buffer.c
}

func (ec *execContext) renameBuffer_autogen(newname, unique lispObject) (lispObject, error) {
	return ec.stub("rename-buffer") // Source file: buffer.c
}

func (ec *execContext) restoreBufferModifiedP_autogen(flag lispObject) (lispObject, error) {
	return ec.stub("restore-buffer-modified-p") // Source file: buffer.c
}

func (ec *execContext) setBuffer_autogen(buffer_or_name lispObject) (lispObject, error) {
	return ec.stub("set-buffer") // Source file: buffer.c
}

func (ec *execContext) setBufferMajorMode_autogen(buffer lispObject) (lispObject, error) {
	return ec.stub("set-buffer-major-mode") // Source file: buffer.c
}

func (ec *execContext) setBufferModifiedP_autogen(flag lispObject) (lispObject, error) {
	return ec.stub("set-buffer-modified-p") // Source file: buffer.c
}

func (ec *execContext) setBufferMultibyte_autogen(flag lispObject) (lispObject, error) {
	return ec.stub("set-buffer-multibyte") // Source file: buffer.c
}

func (ec *execContext) byteCode_autogen(bytestr, vector, maxdepth lispObject) (lispObject, error) {
	return ec.stub("byte-code") // Source file: bytecode.c
}

func (ec *execContext) internalStackStats_autogen() (lispObject, error) {
	return ec.stub("internal-stack-stats") // Source file: bytecode.c
}

func (ec *execContext) callInteractively_autogen(function, record_flag, keys lispObject) (lispObject, error) {
	return ec.stub("call-interactively") // Source file: callint.c
}

func (ec *execContext) funcallInteractively_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("funcall-interactively") // Source file: callint.c
}

func (ec *execContext) interactive_autogen(args lispObject) (lispObject, error) {
	return ec.stub("interactive") // Source file: callint.c, attributes: const
}

func (ec *execContext) prefixNumericValue_autogen(raw lispObject) (lispObject, error) {
	return ec.stub("prefix-numeric-value") // Source file: callint.c
}

func (ec *execContext) callProcess_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("call-process") // Source file: callproc.c
}

func (ec *execContext) callProcessRegion_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("call-process-region") // Source file: callproc.c
}

func (ec *execContext) getenvInternal_autogen(variable, env lispObject) (lispObject, error) {
	return ec.stub("getenv-internal") // Source file: callproc.c
}

func (ec *execContext) capitalize_autogen(obj lispObject) (lispObject, error) {
	return ec.stub("capitalize") // Source file: casefiddle.c
}

func (ec *execContext) capitalizeRegion_autogen(beg, end, region_noncontiguous_p lispObject) (lispObject, error) {
	return ec.stub("capitalize-region") // Source file: casefiddle.c
}

func (ec *execContext) capitalizeWord_autogen(arg lispObject) (lispObject, error) {
	return ec.stub("capitalize-word") // Source file: casefiddle.c
}

func (ec *execContext) downcase_autogen(obj lispObject) (lispObject, error) {
	return ec.stub("downcase") // Source file: casefiddle.c
}

func (ec *execContext) downcaseRegion_autogen(beg, end, region_noncontiguous_p lispObject) (lispObject, error) {
	return ec.stub("downcase-region") // Source file: casefiddle.c
}

func (ec *execContext) downcaseWord_autogen(arg lispObject) (lispObject, error) {
	return ec.stub("downcase-word") // Source file: casefiddle.c
}

func (ec *execContext) upcase_autogen(obj lispObject) (lispObject, error) {
	return ec.stub("upcase") // Source file: casefiddle.c
}

func (ec *execContext) upcaseInitials_autogen(obj lispObject) (lispObject, error) {
	return ec.stub("upcase-initials") // Source file: casefiddle.c
}

func (ec *execContext) upcaseInitialsRegion_autogen(beg, end, region_noncontiguous_p lispObject) (lispObject, error) {
	return ec.stub("upcase-initials-region") // Source file: casefiddle.c
}

func (ec *execContext) upcaseRegion_autogen(beg, end, region_noncontiguous_p lispObject) (lispObject, error) {
	return ec.stub("upcase-region") // Source file: casefiddle.c
}

func (ec *execContext) upcaseWord_autogen(arg lispObject) (lispObject, error) {
	return ec.stub("upcase-word") // Source file: casefiddle.c
}

func (ec *execContext) caseTableP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("case-table-p") // Source file: casetab.c
}

func (ec *execContext) currentCaseTable_autogen() (lispObject, error) {
	return ec.stub("current-case-table") // Source file: casetab.c
}

func (ec *execContext) setCaseTable_autogen(table lispObject) (lispObject, error) {
	return ec.stub("set-case-table") // Source file: casetab.c
}

func (ec *execContext) setStandardCaseTable_autogen(table lispObject) (lispObject, error) {
	return ec.stub("set-standard-case-table") // Source file: casetab.c
}

func (ec *execContext) standardCaseTable_autogen() (lispObject, error) {
	return ec.stub("standard-case-table") // Source file: casetab.c
}

func (ec *execContext) categoryDocstring_autogen(category, table lispObject) (lispObject, error) {
	return ec.stub("category-docstring") // Source file: category.c
}

func (ec *execContext) categorySetMnemonics_autogen(category_set lispObject) (lispObject, error) {
	return ec.stub("category-set-mnemonics") // Source file: category.c
}

func (ec *execContext) categoryTable_autogen() (lispObject, error) {
	return ec.stub("category-table") // Source file: category.c
}

func (ec *execContext) categoryTableP_autogen(arg lispObject) (lispObject, error) {
	return ec.stub("category-table-p") // Source file: category.c
}

func (ec *execContext) charCategorySet_autogen(ch lispObject) (lispObject, error) {
	return ec.stub("char-category-set") // Source file: category.c
}

func (ec *execContext) copyCategoryTable_autogen(table lispObject) (lispObject, error) {
	return ec.stub("copy-category-table") // Source file: category.c
}

func (ec *execContext) defineCategory_autogen(category, docstring, table lispObject) (lispObject, error) {
	return ec.stub("define-category") // Source file: category.c
}

func (ec *execContext) getUnusedCategory_autogen(table lispObject) (lispObject, error) {
	return ec.stub("get-unused-category") // Source file: category.c
}

func (ec *execContext) makeCategorySet_autogen(categories lispObject) (lispObject, error) {
	return ec.stub("make-category-set") // Source file: category.c
}

func (ec *execContext) makeCategoryTable_autogen() (lispObject, error) {
	return ec.stub("make-category-table") // Source file: category.c
}

func (ec *execContext) modifyCategoryEntry_autogen(character, category, table, reset lispObject) (lispObject, error) {
	return ec.stub("modify-category-entry") // Source file: category.c
}

func (ec *execContext) setCategoryTable_autogen(table lispObject) (lispObject, error) {
	return ec.stub("set-category-table") // Source file: category.c
}

func (ec *execContext) standardCategoryTable_autogen() (lispObject, error) {
	return ec.stub("standard-category-table") // Source file: category.c
}

func (ec *execContext) cclExecute_autogen(ccl_prog, reg lispObject) (lispObject, error) {
	return ec.stub("ccl-execute") // Source file: ccl.c
}

func (ec *execContext) cclExecuteOnString_autogen(ccl_prog, status, str, contin, unibyte_p lispObject) (lispObject, error) {
	return ec.stub("ccl-execute-on-string") // Source file: ccl.c
}

func (ec *execContext) cclProgramP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("ccl-program-p") // Source file: ccl.c
}

func (ec *execContext) registerCclProgram_autogen(name, ccl_prog lispObject) (lispObject, error) {
	return ec.stub("register-ccl-program") // Source file: ccl.c
}

func (ec *execContext) registerCodeConversionMap_autogen(symbol, map_ lispObject) (lispObject, error) {
	return ec.stub("register-code-conversion-map") // Source file: ccl.c
}

func (ec *execContext) charResolveModifiers_autogen(character lispObject) (lispObject, error) {
	return ec.stub("char-resolve-modifiers") // Source file: character.c
}

func (ec *execContext) charWidth_autogen(ch lispObject) (lispObject, error) {
	return ec.stub("char-width") // Source file: character.c
}

func (ec *execContext) characterp_autogen(object, ignore lispObject) (lispObject, error) {
	return ec.stub("characterp") // Source file: character.c, attributes: const
}

func (ec *execContext) getByte_autogen(position, string lispObject) (lispObject, error) {
	return ec.stub("get-byte") // Source file: character.c
}

func (ec *execContext) maxChar_autogen(unicode lispObject) (lispObject, error) {
	return ec.stub("max-char") // Source file: character.c, attributes: const
}

func (ec *execContext) multibyteCharToUnibyte_autogen(ch lispObject) (lispObject, error) {
	return ec.stub("multibyte-char-to-unibyte") // Source file: character.c
}

func (ec *execContext) string_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("string") // Source file: character.c
}

func (ec *execContext) stringWidth_autogen(str, from, to lispObject) (lispObject, error) {
	return ec.stub("string-width") // Source file: character.c
}

func (ec *execContext) unibyteCharToMultibyte_autogen(ch lispObject) (lispObject, error) {
	return ec.stub("unibyte-char-to-multibyte") // Source file: character.c
}

func (ec *execContext) unibyteString_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("unibyte-string") // Source file: character.c
}

func (ec *execContext) charCharset_autogen(ch, restriction lispObject) (lispObject, error) {
	return ec.stub("char-charset") // Source file: charset.c
}

func (ec *execContext) charsetAfter_autogen(pos lispObject) (lispObject, error) {
	return ec.stub("charset-after") // Source file: charset.c
}

func (ec *execContext) charsetIdInternal_autogen(charset lispObject) (lispObject, error) {
	return ec.stub("charset-id-internal") // Source file: charset.c
}

func (ec *execContext) charsetPlist_autogen(charset lispObject) (lispObject, error) {
	return ec.stub("charset-plist") // Source file: charset.c
}

func (ec *execContext) charsetPriorityList_autogen(highestp lispObject) (lispObject, error) {
	return ec.stub("charset-priority-list") // Source file: charset.c
}

func (ec *execContext) charsetp_autogen(object lispObject) (lispObject, error) {
	return ec.stub("charsetp") // Source file: charset.c
}

func (ec *execContext) clearCharsetMaps_autogen() (lispObject, error) {
	return ec.stub("clear-charset-maps") // Source file: charset.c
}

func (ec *execContext) declareEquivCharset_autogen(dimension, chars, final_char, charset lispObject) (lispObject, error) {
	return ec.stub("declare-equiv-charset") // Source file: charset.c
}

func (ec *execContext) decodeChar_autogen(charset, code_point lispObject) (lispObject, error) {
	return ec.stub("decode-char") // Source file: charset.c
}

func (ec *execContext) defineCharsetAlias_autogen(alias, charset lispObject) (lispObject, error) {
	return ec.stub("define-charset-alias") // Source file: charset.c
}

func (ec *execContext) defineCharsetInternal_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("define-charset-internal") // Source file: charset.c
}

func (ec *execContext) encodeChar_autogen(ch, charset lispObject) (lispObject, error) {
	return ec.stub("encode-char") // Source file: charset.c
}

func (ec *execContext) findCharsetRegion_autogen(beg, end, table lispObject) (lispObject, error) {
	return ec.stub("find-charset-region") // Source file: charset.c
}

func (ec *execContext) findCharsetString_autogen(str, table lispObject) (lispObject, error) {
	return ec.stub("find-charset-string") // Source file: charset.c
}

func (ec *execContext) getUnusedIsoFinalChar_autogen(dimension, chars lispObject) (lispObject, error) {
	return ec.stub("get-unused-iso-final-char") // Source file: charset.c
}

func (ec *execContext) isoCharset_autogen(dimension, chars, final_char lispObject) (lispObject, error) {
	return ec.stub("iso-charset") // Source file: charset.c
}

func (ec *execContext) makeChar_autogen(charset, code1, code2, code3, code4 lispObject) (lispObject, error) {
	return ec.stub("make-char") // Source file: charset.c
}

func (ec *execContext) mapCharsetChars_autogen(function, charset, arg, from_code, to_code lispObject) (lispObject, error) {
	return ec.stub("map-charset-chars") // Source file: charset.c
}

func (ec *execContext) setCharsetPlist_autogen(charset, plist lispObject) (lispObject, error) {
	return ec.stub("set-charset-plist") // Source file: charset.c
}

func (ec *execContext) setCharsetPriority_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("set-charset-priority") // Source file: charset.c
}

func (ec *execContext) sortCharsets_autogen(charsets lispObject) (lispObject, error) {
	return ec.stub("sort-charsets") // Source file: charset.c
}

func (ec *execContext) splitChar_autogen(ch lispObject) (lispObject, error) {
	return ec.stub("split-char") // Source file: charset.c
}

func (ec *execContext) unifyCharset_autogen(charset, unify_map, deunify lispObject) (lispObject, error) {
	return ec.stub("unify-charset") // Source file: charset.c
}

func (ec *execContext) charTableExtraSlot_autogen(char_table, n lispObject) (lispObject, error) {
	return ec.stub("char-table-extra-slot") // Source file: chartab.c
}

func (ec *execContext) charTableParent_autogen(char_table lispObject) (lispObject, error) {
	return ec.stub("char-table-parent") // Source file: chartab.c
}

func (ec *execContext) charTableRange_autogen(char_table, range_ lispObject) (lispObject, error) {
	return ec.stub("char-table-range") // Source file: chartab.c
}

func (ec *execContext) charTableSubtype_autogen(char_table lispObject) (lispObject, error) {
	return ec.stub("char-table-subtype") // Source file: chartab.c
}

func (ec *execContext) getUnicodePropertyInternal_autogen(char_table, ch lispObject) (lispObject, error) {
	return ec.stub("get-unicode-property-internal") // Source file: chartab.c
}

func (ec *execContext) makeCharTable_autogen(purpose, init lispObject) (lispObject, error) {
	return ec.stub("make-char-table") // Source file: chartab.c
}

func (ec *execContext) mapCharTable_autogen(function, char_table lispObject) (lispObject, error) {
	return ec.stub("map-char-table") // Source file: chartab.c
}

func (ec *execContext) optimizeCharTable_autogen(char_table, test lispObject) (lispObject, error) {
	return ec.stub("optimize-char-table") // Source file: chartab.c
}

func (ec *execContext) putUnicodePropertyInternal_autogen(char_table, ch, value lispObject) (lispObject, error) {
	return ec.stub("put-unicode-property-internal") // Source file: chartab.c
}

func (ec *execContext) setCharTableExtraSlot_autogen(char_table, n, value lispObject) (lispObject, error) {
	return ec.stub("set-char-table-extra-slot") // Source file: chartab.c
}

func (ec *execContext) setCharTableParent_autogen(char_table, parent lispObject) (lispObject, error) {
	return ec.stub("set-char-table-parent") // Source file: chartab.c
}

func (ec *execContext) setCharTableRange_autogen(char_table, range_, value lispObject) (lispObject, error) {
	return ec.stub("set-char-table-range") // Source file: chartab.c
}

func (ec *execContext) unicodePropertyTableInternal_autogen(prop lispObject) (lispObject, error) {
	return ec.stub("unicode-property-table-internal") // Source file: chartab.c
}

func (ec *execContext) backwardChar_autogen(n lispObject) (lispObject, error) {
	return ec.stub("backward-char") // Source file: cmds.c
}

func (ec *execContext) beginningOfLine_autogen(n lispObject) (lispObject, error) {
	return ec.stub("beginning-of-line") // Source file: cmds.c
}

func (ec *execContext) deleteChar_autogen(n, killflag lispObject) (lispObject, error) {
	return ec.stub("delete-char") // Source file: cmds.c
}

func (ec *execContext) endOfLine_autogen(n lispObject) (lispObject, error) {
	return ec.stub("end-of-line") // Source file: cmds.c
}

func (ec *execContext) forwardChar_autogen(n lispObject) (lispObject, error) {
	return ec.stub("forward-char") // Source file: cmds.c
}

func (ec *execContext) forwardLine_autogen(n lispObject) (lispObject, error) {
	return ec.stub("forward-line") // Source file: cmds.c
}

func (ec *execContext) selfInsertCommand_autogen(n, c lispObject) (lispObject, error) {
	return ec.stub("self-insert-command") // Source file: cmds.c
}

func (ec *execContext) checkCodingSystem_autogen(coding_system lispObject) (lispObject, error) {
	return ec.stub("check-coding-system") // Source file: coding.c
}

func (ec *execContext) checkCodingSystemsRegion_autogen(start, end, coding_system_list lispObject) (lispObject, error) {
	return ec.stub("check-coding-systems-region") // Source file: coding.c
}

func (ec *execContext) codingSystemAliases_autogen(coding_system lispObject) (lispObject, error) {
	return ec.stub("coding-system-aliases") // Source file: coding.c
}

func (ec *execContext) codingSystemBase_autogen(coding_system lispObject) (lispObject, error) {
	return ec.stub("coding-system-base") // Source file: coding.c
}

func (ec *execContext) codingSystemEolType_autogen(coding_system lispObject) (lispObject, error) {
	return ec.stub("coding-system-eol-type") // Source file: coding.c
}

func (ec *execContext) codingSystemP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("coding-system-p") // Source file: coding.c
}

func (ec *execContext) codingSystemPlist_autogen(coding_system lispObject) (lispObject, error) {
	return ec.stub("coding-system-plist") // Source file: coding.c
}

func (ec *execContext) codingSystemPriorityList_autogen(highestp lispObject) (lispObject, error) {
	return ec.stub("coding-system-priority-list") // Source file: coding.c
}

func (ec *execContext) codingSystemPut_autogen(coding_system, prop, val lispObject) (lispObject, error) {
	return ec.stub("coding-system-put") // Source file: coding.c
}

func (ec *execContext) decodeBig5Char_autogen(code lispObject) (lispObject, error) {
	return ec.stub("decode-big5-char") // Source file: coding.c
}

func (ec *execContext) decodeCodingRegion_autogen(start, end, coding_system, destination lispObject) (lispObject, error) {
	return ec.stub("decode-coding-region") // Source file: coding.c
}

func (ec *execContext) decodeCodingString_autogen(string, coding_system, nocopy, buffer lispObject) (lispObject, error) {
	return ec.stub("decode-coding-string") // Source file: coding.c
}

func (ec *execContext) decodeSjisChar_autogen(code lispObject) (lispObject, error) {
	return ec.stub("decode-sjis-char") // Source file: coding.c
}

func (ec *execContext) defineCodingSystemAlias_autogen(alias, coding_system lispObject) (lispObject, error) {
	return ec.stub("define-coding-system-alias") // Source file: coding.c
}

func (ec *execContext) defineCodingSystemInternal_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("define-coding-system-internal") // Source file: coding.c
}

func (ec *execContext) detectCodingRegion_autogen(start, end, highest lispObject) (lispObject, error) {
	return ec.stub("detect-coding-region") // Source file: coding.c
}

func (ec *execContext) detectCodingString_autogen(string, highest lispObject) (lispObject, error) {
	return ec.stub("detect-coding-string") // Source file: coding.c
}

func (ec *execContext) encodeBig5Char_autogen(ch lispObject) (lispObject, error) {
	return ec.stub("encode-big5-char") // Source file: coding.c
}

func (ec *execContext) encodeCodingRegion_autogen(start, end, coding_system, destination lispObject) (lispObject, error) {
	return ec.stub("encode-coding-region") // Source file: coding.c
}

func (ec *execContext) encodeCodingString_autogen(string, coding_system, nocopy, buffer lispObject) (lispObject, error) {
	return ec.stub("encode-coding-string") // Source file: coding.c
}

func (ec *execContext) encodeSjisChar_autogen(ch lispObject) (lispObject, error) {
	return ec.stub("encode-sjis-char") // Source file: coding.c
}

func (ec *execContext) findCodingSystemsRegionInternal_autogen(start, end, exclude lispObject) (lispObject, error) {
	return ec.stub("find-coding-systems-region-internal") // Source file: coding.c
}

func (ec *execContext) findOperationCodingSystem_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("find-operation-coding-system") // Source file: coding.c
}

func (ec *execContext) internalDecodeStringUtf8_autogen(string, buffer, nocopy, handle_8_bit, handle_over_uni, decode_method, count lispObject) (lispObject, error) {
	return ec.stub("internal-decode-string-utf-8") // Source file: coding.c
}

func (ec *execContext) internalEncodeStringUtf8_autogen(string, buffer, nocopy, handle_8_bit, handle_over_uni, encode_method, count lispObject) (lispObject, error) {
	return ec.stub("internal-encode-string-utf-8") // Source file: coding.c
}

func (ec *execContext) keyboardCodingSystem_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("keyboard-coding-system") // Source file: coding.c
}

func (ec *execContext) readCodingSystem_autogen(prompt, default_coding_system lispObject) (lispObject, error) {
	return ec.stub("read-coding-system") // Source file: coding.c
}

func (ec *execContext) readNonNilCodingSystem_autogen(prompt lispObject) (lispObject, error) {
	return ec.stub("read-non-nil-coding-system") // Source file: coding.c
}

func (ec *execContext) setCodingSystemPriority_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("set-coding-system-priority") // Source file: coding.c
}

func (ec *execContext) setKeyboardCodingSystemInternal_autogen(coding_system, terminal lispObject) (lispObject, error) {
	return ec.stub("set-keyboard-coding-system-internal") // Source file: coding.c
}

func (ec *execContext) setSafeTerminalCodingSystemInternal_autogen(coding_system lispObject) (lispObject, error) {
	return ec.stub("set-safe-terminal-coding-system-internal") // Source file: coding.c
}

func (ec *execContext) setTerminalCodingSystemInternal_autogen(coding_system, terminal lispObject) (lispObject, error) {
	return ec.stub("set-terminal-coding-system-internal") // Source file: coding.c
}

func (ec *execContext) terminalCodingSystem_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("terminal-coding-system") // Source file: coding.c
}

func (ec *execContext) unencodableCharPosition_autogen(start, end, coding_system, count, string lispObject) (lispObject, error) {
	return ec.stub("unencodable-char-position") // Source file: coding.c
}

func (ec *execContext) compCompileCtxtToFile_autogen(filename lispObject) (lispObject, error) {
	return ec.stub("comp--compile-ctxt-to-file") // Source file: comp.c
}

func (ec *execContext) compInitCtxt_autogen() (lispObject, error) {
	return ec.stub("comp--init-ctxt") // Source file: comp.c
}

func (ec *execContext) compInstallTrampoline_autogen(subr_name, trampoline lispObject) (lispObject, error) {
	return ec.stub("comp--install-trampoline") // Source file: comp.c
}

func (ec *execContext) compLateRegisterSubr_autogen(name, c_name, minarg, maxarg, type_, rest, comp_u lispObject) (lispObject, error) {
	return ec.stub("comp--late-register-subr") // Source file: comp.c
}

func (ec *execContext) compRegisterLambda_autogen(reloc_idx, c_name, minarg, maxarg, type_, rest, comp_u lispObject) (lispObject, error) {
	return ec.stub("comp--register-lambda") // Source file: comp.c
}

func (ec *execContext) compRegisterSubr_autogen(name, c_name, minarg, maxarg, type_, rest, comp_u lispObject) (lispObject, error) {
	return ec.stub("comp--register-subr") // Source file: comp.c
}

func (ec *execContext) compReleaseCtxt_autogen() (lispObject, error) {
	return ec.stub("comp--release-ctxt") // Source file: comp.c
}

func (ec *execContext) compSubrSignature_autogen(subr lispObject) (lispObject, error) {
	return ec.stub("comp--subr-signature") // Source file: comp.c
}

func (ec *execContext) compElToElnFilename_autogen(filename, base_dir lispObject) (lispObject, error) {
	return ec.stub("comp-el-to-eln-filename") // Source file: comp.c
}

func (ec *execContext) compElToElnRelFilename_autogen(filename lispObject) (lispObject, error) {
	return ec.stub("comp-el-to-eln-rel-filename") // Source file: comp.c
}

func (ec *execContext) compLibgccjitVersion_autogen() (lispObject, error) {
	return ec.stub("comp-libgccjit-version") // Source file: comp.c
}

func (ec *execContext) compNativeCompilerOptionsEffectiveP_autogen() (lispObject, error) {
	return ec.stub("comp-native-compiler-options-effective-p") // Source file: comp.c
}

func (ec *execContext) compNativeDriverOptionsEffectiveP_autogen() (lispObject, error) {
	return ec.stub("comp-native-driver-options-effective-p") // Source file: comp.c
}

func (ec *execContext) nativeCompAvailableP_autogen() (lispObject, error) {
	return ec.stub("native-comp-available-p") // Source file: comp.c
}

func (ec *execContext) nativeElispLoad_autogen(filename, late_load lispObject) (lispObject, error) {
	return ec.stub("native-elisp-load") // Source file: comp.c
}

func (ec *execContext) clearCompositionCache_autogen() (lispObject, error) {
	return ec.stub("clear-composition-cache") // Source file: composite.c
}

func (ec *execContext) composeRegionInternal_autogen(start, end, components, modification_func lispObject) (lispObject, error) {
	return ec.stub("compose-region-internal") // Source file: composite.c
}

func (ec *execContext) composeStringInternal_autogen(string, start, end, components, modification_func lispObject) (lispObject, error) {
	return ec.stub("compose-string-internal") // Source file: composite.c
}

func (ec *execContext) compositionGetGstring_autogen(from, to, font_object, string lispObject) (lispObject, error) {
	return ec.stub("composition-get-gstring") // Source file: composite.c
}

func (ec *execContext) compositionSortRules_autogen(rules lispObject) (lispObject, error) {
	return ec.stub("composition-sort-rules") // Source file: composite.c
}

func (ec *execContext) findCompositionInternal_autogen(pos, limit, string, detail_p lispObject) (lispObject, error) {
	return ec.stub("find-composition-internal") // Source file: composite.c
}

func (ec *execContext) cygwinConvertFileNameFromWindows_autogen(file, absolute_p lispObject) (lispObject, error) {
	return ec.stub("cygwin-convert-file-name-from-windows") // Source file: cygw32.c
}

func (ec *execContext) cygwinConvertFileNameToWindows_autogen(file, absolute_p lispObject) (lispObject, error) {
	return ec.stub("cygwin-convert-file-name-to-windows") // Source file: cygw32.c
}

func (ec *execContext) rem_autogen(x, y lispObject) (lispObject, error) {
	return ec.stub("%") // Source file: data.c
}

func (ec *execContext) times_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("*") // Source file: data.c
}

func (ec *execContext) plus_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("+") // Source file: data.c
}

func (ec *execContext) minus_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("-") // Source file: data.c
}

func (ec *execContext) quo_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("/") // Source file: data.c
}

func (ec *execContext) neq_autogen(num1, num2 lispObject) (lispObject, error) {
	return ec.stub("/=") // Source file: data.c
}

func (ec *execContext) add1_autogen(number lispObject) (lispObject, error) {
	return ec.stub("1+") // Source file: data.c
}

func (ec *execContext) sub1_autogen(number lispObject) (lispObject, error) {
	return ec.stub("1-") // Source file: data.c
}

func (ec *execContext) lss_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("<") // Source file: data.c
}

func (ec *execContext) leq_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("<=") // Source file: data.c
}

func (ec *execContext) eqlsign_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("=") // Source file: data.c
}

func (ec *execContext) gtr_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub(">") // Source file: data.c
}

func (ec *execContext) geq_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub(">=") // Source file: data.c
}

func (ec *execContext) addVariableWatcher_autogen(symbol, watch_function lispObject) (lispObject, error) {
	return ec.stub("add-variable-watcher") // Source file: data.c
}

func (ec *execContext) aref_autogen(array, idx lispObject) (lispObject, error) {
	return ec.stub("aref") // Source file: data.c
}

func (ec *execContext) arrayp_autogen(object lispObject) (lispObject, error) {
	return ec.stub("arrayp") // Source file: data.c
}

func (ec *execContext) aset_autogen(array, idx, newelt lispObject) (lispObject, error) {
	return ec.stub("aset") // Source file: data.c
}

func (ec *execContext) ash_autogen(value, count lispObject) (lispObject, error) {
	return ec.stub("ash") // Source file: data.c
}

func (ec *execContext) atom_autogen(object lispObject) (lispObject, error) {
	return ec.stub("atom") // Source file: data.c, attributes: const
}

func (ec *execContext) bareSymbol_autogen(sym lispObject) (lispObject, error) {
	return ec.stub("bare-symbol") // Source file: data.c
}

func (ec *execContext) bareSymbolP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("bare-symbol-p") // Source file: data.c, attributes: const
}

func (ec *execContext) boolVectorCountConsecutive_autogen(a, b, i lispObject) (lispObject, error) {
	return ec.stub("bool-vector-count-consecutive") // Source file: data.c
}

func (ec *execContext) boolVectorCountPopulation_autogen(a lispObject) (lispObject, error) {
	return ec.stub("bool-vector-count-population") // Source file: data.c
}

func (ec *execContext) boolVectorExclusiveOr_autogen(a, b, c lispObject) (lispObject, error) {
	return ec.stub("bool-vector-exclusive-or") // Source file: data.c
}

func (ec *execContext) boolVectorIntersection_autogen(a, b, c lispObject) (lispObject, error) {
	return ec.stub("bool-vector-intersection") // Source file: data.c
}

func (ec *execContext) boolVectorNot_autogen(a, b lispObject) (lispObject, error) {
	return ec.stub("bool-vector-not") // Source file: data.c
}

func (ec *execContext) boolVectorP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("bool-vector-p") // Source file: data.c
}

func (ec *execContext) boolVectorSetDifference_autogen(a, b, c lispObject) (lispObject, error) {
	return ec.stub("bool-vector-set-difference") // Source file: data.c
}

func (ec *execContext) boolVectorSubsetp_autogen(a, b lispObject) (lispObject, error) {
	return ec.stub("bool-vector-subsetp") // Source file: data.c
}

func (ec *execContext) boolVectorUnion_autogen(a, b, c lispObject) (lispObject, error) {
	return ec.stub("bool-vector-union") // Source file: data.c
}

func (ec *execContext) boundp_autogen(symbol lispObject) (lispObject, error) {
	return ec.stub("boundp") // Source file: data.c
}

func (ec *execContext) bufferp_autogen(object lispObject) (lispObject, error) {
	return ec.stub("bufferp") // Source file: data.c
}

func (ec *execContext) byteCodeFunctionP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("byte-code-function-p") // Source file: data.c
}

func (ec *execContext) byteorder_autogen() (lispObject, error) {
	return ec.stub("byteorder") // Source file: data.c, attributes: const
}

func (ec *execContext) car_autogen(list lispObject) (lispObject, error) {
	return ec.stub("car") // Source file: data.c
}

func (ec *execContext) carSafe_autogen(object lispObject) (lispObject, error) {
	return ec.stub("car-safe") // Source file: data.c
}

func (ec *execContext) cdr_autogen(list lispObject) (lispObject, error) {
	return ec.stub("cdr") // Source file: data.c
}

func (ec *execContext) cdrSafe_autogen(object lispObject) (lispObject, error) {
	return ec.stub("cdr-safe") // Source file: data.c
}

func (ec *execContext) charOrStringP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("char-or-string-p") // Source file: data.c, attributes: const
}

func (ec *execContext) charTableP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("char-table-p") // Source file: data.c
}

func (ec *execContext) commandModes_autogen(command lispObject) (lispObject, error) {
	return ec.stub("command-modes") // Source file: data.c
}

func (ec *execContext) conditionVariableP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("condition-variable-p") // Source file: data.c
}

func (ec *execContext) consp_autogen(object lispObject) (lispObject, error) {
	return ec.stub("consp") // Source file: data.c, attributes: const
}

func (ec *execContext) defalias_autogen(symbol, definition, docstring lispObject) (lispObject, error) {
	return ec.stub("defalias") // Source file: data.c
}

func (ec *execContext) defaultBoundp_autogen(symbol lispObject) (lispObject, error) {
	return ec.stub("default-boundp") // Source file: data.c
}

func (ec *execContext) defaultValue_autogen(symbol lispObject) (lispObject, error) {
	return ec.stub("default-value") // Source file: data.c
}

func (ec *execContext) eq_autogen(obj1, obj2 lispObject) (lispObject, error) {
	return ec.stub("eq") // Source file: data.c, attributes: const
}

func (ec *execContext) fboundp_autogen(symbol lispObject) (lispObject, error) {
	return ec.stub("fboundp") // Source file: data.c
}

func (ec *execContext) floatp_autogen(object lispObject) (lispObject, error) {
	return ec.stub("floatp") // Source file: data.c, attributes: const
}

func (ec *execContext) fmakunbound_autogen(symbol lispObject) (lispObject, error) {
	return ec.stub("fmakunbound") // Source file: data.c
}

func (ec *execContext) fset_autogen(symbol, definition lispObject) (lispObject, error) {
	return ec.stub("fset") // Source file: data.c
}

func (ec *execContext) getVariableWatchers_autogen(symbol lispObject) (lispObject, error) {
	return ec.stub("get-variable-watchers") // Source file: data.c
}

func (ec *execContext) indirectFunction_autogen(object, noerror lispObject) (lispObject, error) {
	return ec.stub("indirect-function") // Source file: data.c
}

func (ec *execContext) indirectVariable_autogen(object lispObject) (lispObject, error) {
	return ec.stub("indirect-variable") // Source file: data.c
}

func (ec *execContext) integerOrMarkerP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("integer-or-marker-p") // Source file: data.c
}

func (ec *execContext) integerp_autogen(object lispObject) (lispObject, error) {
	return ec.stub("integerp") // Source file: data.c, attributes: const
}

func (ec *execContext) interactiveForm_autogen(cmd lispObject) (lispObject, error) {
	return ec.stub("interactive-form") // Source file: data.c
}

func (ec *execContext) keywordp_autogen(object lispObject) (lispObject, error) {
	return ec.stub("keywordp") // Source file: data.c
}

func (ec *execContext) killLocalVariable_autogen(variable lispObject) (lispObject, error) {
	return ec.stub("kill-local-variable") // Source file: data.c
}

func (ec *execContext) listp_autogen(object lispObject) (lispObject, error) {
	return ec.stub("listp") // Source file: data.c, attributes: const
}

func (ec *execContext) localVariableIfSetP_autogen(variable, buffer lispObject) (lispObject, error) {
	return ec.stub("local-variable-if-set-p") // Source file: data.c
}

func (ec *execContext) localVariableP_autogen(variable, buffer lispObject) (lispObject, error) {
	return ec.stub("local-variable-p") // Source file: data.c
}

func (ec *execContext) logand_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("logand") // Source file: data.c
}

func (ec *execContext) logcount_autogen(value lispObject) (lispObject, error) {
	return ec.stub("logcount") // Source file: data.c
}

func (ec *execContext) logior_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("logior") // Source file: data.c
}

func (ec *execContext) lognot_autogen(number lispObject) (lispObject, error) {
	return ec.stub("lognot") // Source file: data.c
}

func (ec *execContext) logxor_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("logxor") // Source file: data.c
}

func (ec *execContext) makeLocalVariable_autogen(variable lispObject) (lispObject, error) {
	return ec.stub("make-local-variable") // Source file: data.c
}

func (ec *execContext) makeVariableBufferLocal_autogen(variable lispObject) (lispObject, error) {
	return ec.stub("make-variable-buffer-local") // Source file: data.c
}

func (ec *execContext) makunbound_autogen(symbol lispObject) (lispObject, error) {
	return ec.stub("makunbound") // Source file: data.c
}

func (ec *execContext) markerp_autogen(object lispObject) (lispObject, error) {
	return ec.stub("markerp") // Source file: data.c
}

func (ec *execContext) max_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("max") // Source file: data.c
}

func (ec *execContext) min_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("min") // Source file: data.c
}

func (ec *execContext) mod_autogen(x, y lispObject) (lispObject, error) {
	return ec.stub("mod") // Source file: data.c
}

func (ec *execContext) moduleFunctionP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("module-function-p") // Source file: data.c, attributes: const
}

func (ec *execContext) multibyteStringP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("multibyte-string-p") // Source file: data.c
}

func (ec *execContext) mutexp_autogen(object lispObject) (lispObject, error) {
	return ec.stub("mutexp") // Source file: data.c
}

func (ec *execContext) nativeCompUnitFile_autogen(comp_unit lispObject) (lispObject, error) {
	return ec.stub("native-comp-unit-file") // Source file: data.c
}

func (ec *execContext) nativeCompUnitSetFile_autogen(comp_unit, new_file lispObject) (lispObject, error) {
	return ec.stub("native-comp-unit-set-file") // Source file: data.c
}

func (ec *execContext) natnump_autogen(object lispObject) (lispObject, error) {
	return ec.stub("natnump") // Source file: data.c, attributes: const
}

func (ec *execContext) nlistp_autogen(object lispObject) (lispObject, error) {
	return ec.stub("nlistp") // Source file: data.c, attributes: const
}

func (ec *execContext) null_autogen(object lispObject) (lispObject, error) {
	return ec.stub("null") // Source file: data.c, attributes: const
}

func (ec *execContext) numberOrMarkerP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("number-or-marker-p") // Source file: data.c
}

func (ec *execContext) numberToString_autogen(number lispObject) (lispObject, error) {
	return ec.stub("number-to-string") // Source file: data.c
}

func (ec *execContext) numberp_autogen(object lispObject) (lispObject, error) {
	return ec.stub("numberp") // Source file: data.c, attributes: const
}

func (ec *execContext) positionSymbol_autogen(sym, pos lispObject) (lispObject, error) {
	return ec.stub("position-symbol") // Source file: data.c
}

func (ec *execContext) recordp_autogen(object lispObject) (lispObject, error) {
	return ec.stub("recordp") // Source file: data.c
}

func (ec *execContext) removePosFromSymbol_autogen(arg lispObject) (lispObject, error) {
	return ec.stub("remove-pos-from-symbol") // Source file: data.c
}

func (ec *execContext) removeVariableWatcher_autogen(symbol, watch_function lispObject) (lispObject, error) {
	return ec.stub("remove-variable-watcher") // Source file: data.c
}

func (ec *execContext) sequencep_autogen(object lispObject) (lispObject, error) {
	return ec.stub("sequencep") // Source file: data.c
}

func (ec *execContext) set_autogen(symbol, newval lispObject) (lispObject, error) {
	return ec.stub("set") // Source file: data.c
}

func (ec *execContext) setDefault_autogen(symbol, value lispObject) (lispObject, error) {
	return ec.stub("set-default") // Source file: data.c
}

func (ec *execContext) setcar_autogen(cell, newcar lispObject) (lispObject, error) {
	return ec.stub("setcar") // Source file: data.c
}

func (ec *execContext) setcdr_autogen(cell, newcdr lispObject) (lispObject, error) {
	return ec.stub("setcdr") // Source file: data.c
}

func (ec *execContext) setplist_autogen(symbol, newplist lispObject) (lispObject, error) {
	return ec.stub("setplist") // Source file: data.c
}

func (ec *execContext) stringToNumber_autogen(string, base lispObject) (lispObject, error) {
	return ec.stub("string-to-number") // Source file: data.c
}

func (ec *execContext) stringp_autogen(object lispObject) (lispObject, error) {
	return ec.stub("stringp") // Source file: data.c, attributes: const
}

func (ec *execContext) subrArity_autogen(subr lispObject) (lispObject, error) {
	return ec.stub("subr-arity") // Source file: data.c
}

func (ec *execContext) subrName_autogen(subr lispObject) (lispObject, error) {
	return ec.stub("subr-name") // Source file: data.c
}

func (ec *execContext) subrNativeCompUnit_autogen(subr lispObject) (lispObject, error) {
	return ec.stub("subr-native-comp-unit") // Source file: data.c
}

func (ec *execContext) subrNativeElispP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("subr-native-elisp-p") // Source file: data.c
}

func (ec *execContext) subrNativeLambdaList_autogen(subr lispObject) (lispObject, error) {
	return ec.stub("subr-native-lambda-list") // Source file: data.c
}

func (ec *execContext) subrType_autogen(subr lispObject) (lispObject, error) {
	return ec.stub("subr-type") // Source file: data.c
}

func (ec *execContext) subrp_autogen(object lispObject) (lispObject, error) {
	return ec.stub("subrp") // Source file: data.c
}

func (ec *execContext) symbolFunction_autogen(symbol lispObject) (lispObject, error) {
	return ec.stub("symbol-function") // Source file: data.c
}

func (ec *execContext) symbolName_autogen(symbol lispObject) (lispObject, error) {
	return ec.stub("symbol-name") // Source file: data.c
}

func (ec *execContext) symbolPlist_autogen(symbol lispObject) (lispObject, error) {
	return ec.stub("symbol-plist") // Source file: data.c
}

func (ec *execContext) symbolValue_autogen(symbol lispObject) (lispObject, error) {
	return ec.stub("symbol-value") // Source file: data.c
}

func (ec *execContext) symbolWithPosP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("symbol-with-pos-p") // Source file: data.c, attributes: const
}

func (ec *execContext) symbolWithPosPos_autogen(ls lispObject) (lispObject, error) {
	return ec.stub("symbol-with-pos-pos") // Source file: data.c
}

func (ec *execContext) symbolp_autogen(object lispObject) (lispObject, error) {
	return ec.stub("symbolp") // Source file: data.c, attributes: const
}

func (ec *execContext) threadp_autogen(object lispObject) (lispObject, error) {
	return ec.stub("threadp") // Source file: data.c
}

func (ec *execContext) typeOf_autogen(object lispObject) (lispObject, error) {
	return ec.stub("type-of") // Source file: data.c
}

func (ec *execContext) userPtrp_autogen(object lispObject) (lispObject, error) {
	return ec.stub("user-ptrp") // Source file: data.c
}

func (ec *execContext) variableBindingLocus_autogen(variable lispObject) (lispObject, error) {
	return ec.stub("variable-binding-locus") // Source file: data.c
}

func (ec *execContext) vectorOrCharTableP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("vector-or-char-table-p") // Source file: data.c
}

func (ec *execContext) vectorp_autogen(object lispObject) (lispObject, error) {
	return ec.stub("vectorp") // Source file: data.c
}

func (ec *execContext) dbusInitBus_autogen(bus, private lispObject) (lispObject, error) {
	return ec.stub("dbus--init-bus") // Source file: dbusbind.c
}

func (ec *execContext) dbusGetUniqueName_autogen(bus lispObject) (lispObject, error) {
	return ec.stub("dbus-get-unique-name") // Source file: dbusbind.c
}

func (ec *execContext) dbusMessageInternal_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("dbus-message-internal") // Source file: dbusbind.c
}

func (ec *execContext) zlibAvailableP_autogen() (lispObject, error) {
	return ec.stub("zlib-available-p") // Source file: decompress.c
}

func (ec *execContext) zlibDecompressRegion_autogen(start, end, allow_partial lispObject) (lispObject, error) {
	return ec.stub("zlib-decompress-region") // Source file: decompress.c
}

func (ec *execContext) directoryFiles_autogen(directory, full, match, nosort, count lispObject) (lispObject, error) {
	return ec.stub("directory-files") // Source file: dired.c
}

func (ec *execContext) directoryFilesAndAttributes_autogen(directory, full, match, nosort, id_format, count lispObject) (lispObject, error) {
	return ec.stub("directory-files-and-attributes") // Source file: dired.c
}

func (ec *execContext) fileAttributes_autogen(filename, id_format lispObject) (lispObject, error) {
	return ec.stub("file-attributes") // Source file: dired.c
}

func (ec *execContext) fileAttributesLessp_autogen(f1, f2 lispObject) (lispObject, error) {
	return ec.stub("file-attributes-lessp") // Source file: dired.c
}

func (ec *execContext) fileNameAllCompletions_autogen(file, directory lispObject) (lispObject, error) {
	return ec.stub("file-name-all-completions") // Source file: dired.c
}

func (ec *execContext) fileNameCompletion_autogen(file, directory, predicate lispObject) (lispObject, error) {
	return ec.stub("file-name-completion") // Source file: dired.c
}

func (ec *execContext) systemGroups_autogen() (lispObject, error) {
	return ec.stub("system-groups") // Source file: dired.c
}

func (ec *execContext) systemUsers_autogen() (lispObject, error) {
	return ec.stub("system-users") // Source file: dired.c
}

func (ec *execContext) ding_autogen(arg lispObject) (lispObject, error) {
	return ec.stub("ding") // Source file: dispnew.c
}

func (ec *execContext) displayUpdateForMouseMovement_autogen(mouse_x, mouse_y lispObject) (lispObject, error) {
	return ec.stub("display--update-for-mouse-movement") // Source file: dispnew.c
}

func (ec *execContext) dumpRedisplayHistory_autogen() (lispObject, error) {
	return ec.stub("dump-redisplay-history") // Source file: dispnew.c
}

func (ec *execContext) frameOrBufferChangedP_autogen(variable lispObject) (lispObject, error) {
	return ec.stub("frame-or-buffer-changed-p") // Source file: dispnew.c
}

func (ec *execContext) internalShowCursor_autogen(window, show lispObject) (lispObject, error) {
	return ec.stub("internal-show-cursor") // Source file: dispnew.c
}

func (ec *execContext) internalShowCursorP_autogen(window lispObject) (lispObject, error) {
	return ec.stub("internal-show-cursor-p") // Source file: dispnew.c
}

func (ec *execContext) openTermscript_autogen(file lispObject) (lispObject, error) {
	return ec.stub("open-termscript") // Source file: dispnew.c
}

func (ec *execContext) redisplay_autogen(force lispObject) (lispObject, error) {
	return ec.stub("redisplay") // Source file: dispnew.c
}

func (ec *execContext) redrawDisplay_autogen() (lispObject, error) {
	return ec.stub("redraw-display") // Source file: dispnew.c
}

func (ec *execContext) redrawFrame_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("redraw-frame") // Source file: dispnew.c
}

func (ec *execContext) sendStringToTerminal_autogen(string, terminal lispObject) (lispObject, error) {
	return ec.stub("send-string-to-terminal") // Source file: dispnew.c
}

func (ec *execContext) sleepFor_autogen(seconds, milliseconds lispObject) (lispObject, error) {
	return ec.stub("sleep-for") // Source file: dispnew.c
}

func (ec *execContext) snarfDocumentation_autogen(filename lispObject) (lispObject, error) {
	return ec.stub("Snarf-documentation") // Source file: doc.c
}

func (ec *execContext) documentation_autogen(function, raw lispObject) (lispObject, error) {
	return ec.stub("documentation") // Source file: doc.c
}

func (ec *execContext) documentationProperty_autogen(symbol, prop, raw lispObject) (lispObject, error) {
	return ec.stub("documentation-property") // Source file: doc.c
}

func (ec *execContext) textQuotingStyle_autogen() (lispObject, error) {
	return ec.stub("text-quoting-style") // Source file: doc.c
}

func (ec *execContext) fileSystemInfo_autogen(filename lispObject) (lispObject, error) {
	return ec.stub("file-system-info") // Source file: dosfns.c
}

func (ec *execContext) fileSystemInfo_1_autogen(filename lispObject) (lispObject, error) {
	return ec.stub("file-system-info") // Source file: fileio.c
}

func (ec *execContext) fileSystemInfo_2_autogen(filename lispObject) (lispObject, error) {
	return ec.stub("file-system-info") // Source file: w32fns.c
}

func (ec *execContext) insertStartupScreen_autogen() (lispObject, error) {
	return ec.stub("insert-startup-screen") // Source file: dosfns.c
}

func (ec *execContext) int86_autogen(interrupt, registers lispObject) (lispObject, error) {
	return ec.stub("int86") // Source file: dosfns.c
}

func (ec *execContext) dosMemget_autogen(address, vector lispObject) (lispObject, error) {
	return ec.stub("msdos-memget") // Source file: dosfns.c
}

func (ec *execContext) dosMemput_autogen(address, vector lispObject) (lispObject, error) {
	return ec.stub("msdos-memput") // Source file: dosfns.c
}

func (ec *execContext) msdosMouseDisable_autogen() (lispObject, error) {
	return ec.stub("msdos-mouse-disable") // Source file: dosfns.c
}

func (ec *execContext) msdosMouseEnable_autogen() (lispObject, error) {
	return ec.stub("msdos-mouse-enable") // Source file: dosfns.c
}

func (ec *execContext) msdosMouseInit_autogen() (lispObject, error) {
	return ec.stub("msdos-mouse-init") // Source file: dosfns.c
}

func (ec *execContext) msdosMouseP_autogen() (lispObject, error) {
	return ec.stub("msdos-mouse-p") // Source file: dosfns.c
}

func (ec *execContext) msdosSetKeyboard_autogen(country_code, allkeys lispObject) (lispObject, error) {
	return ec.stub("msdos-set-keyboard") // Source file: dosfns.c
}

func (ec *execContext) bobp_autogen() (lispObject, error) {
	return ec.stub("bobp") // Source file: editfns.c
}

func (ec *execContext) bolp_autogen() (lispObject, error) {
	return ec.stub("bolp") // Source file: editfns.c
}

func (ec *execContext) bufferSize_autogen(buffer lispObject) (lispObject, error) {
	return ec.stub("buffer-size") // Source file: editfns.c
}

func (ec *execContext) bufferString_autogen() (lispObject, error) {
	return ec.stub("buffer-string") // Source file: editfns.c
}

func (ec *execContext) bufferSubstring_autogen(start, end lispObject) (lispObject, error) {
	return ec.stub("buffer-substring") // Source file: editfns.c
}

func (ec *execContext) bufferSubstringNoProperties_autogen(start, end lispObject) (lispObject, error) {
	return ec.stub("buffer-substring-no-properties") // Source file: editfns.c
}

func (ec *execContext) byteToPosition_autogen(bytepos lispObject) (lispObject, error) {
	return ec.stub("byte-to-position") // Source file: editfns.c
}

func (ec *execContext) byteToString_autogen(byte lispObject) (lispObject, error) {
	return ec.stub("byte-to-string") // Source file: editfns.c
}

func (ec *execContext) charAfter_autogen(pos lispObject) (lispObject, error) {
	return ec.stub("char-after") // Source file: editfns.c
}

func (ec *execContext) charBefore_autogen(pos lispObject) (lispObject, error) {
	return ec.stub("char-before") // Source file: editfns.c
}

func (ec *execContext) charEqual_autogen(c1, c2 lispObject) (lispObject, error) {
	return ec.stub("char-equal") // Source file: editfns.c
}

func (ec *execContext) charToString_autogen(character lispObject) (lispObject, error) {
	return ec.stub("char-to-string") // Source file: editfns.c
}

func (ec *execContext) compareBufferSubstrings_autogen(buffer1, start1, end1, buffer2, start2, end2 lispObject) (lispObject, error) {
	return ec.stub("compare-buffer-substrings") // Source file: editfns.c
}

func (ec *execContext) constrainToField_autogen(new_pos, old_pos, escape_from_edge, only_in_line, inhibit_capture_property lispObject) (lispObject, error) {
	return ec.stub("constrain-to-field") // Source file: editfns.c
}

func (ec *execContext) currentMessage_autogen() (lispObject, error) {
	return ec.stub("current-message") // Source file: editfns.c
}

func (ec *execContext) deleteAndExtractRegion_autogen(start, end lispObject) (lispObject, error) {
	return ec.stub("delete-and-extract-region") // Source file: editfns.c
}

func (ec *execContext) deleteField_autogen(pos lispObject) (lispObject, error) {
	return ec.stub("delete-field") // Source file: editfns.c
}

func (ec *execContext) deleteRegion_autogen(start, end lispObject) (lispObject, error) {
	return ec.stub("delete-region") // Source file: editfns.c
}

func (ec *execContext) emacsPid_autogen() (lispObject, error) {
	return ec.stub("emacs-pid") // Source file: editfns.c
}

func (ec *execContext) eobp_autogen() (lispObject, error) {
	return ec.stub("eobp") // Source file: editfns.c
}

func (ec *execContext) eolp_autogen() (lispObject, error) {
	return ec.stub("eolp") // Source file: editfns.c
}

func (ec *execContext) fieldBeginning_autogen(pos, escape_from_edge, limit lispObject) (lispObject, error) {
	return ec.stub("field-beginning") // Source file: editfns.c
}

func (ec *execContext) fieldEnd_autogen(pos, escape_from_edge, limit lispObject) (lispObject, error) {
	return ec.stub("field-end") // Source file: editfns.c
}

func (ec *execContext) fieldString_autogen(pos lispObject) (lispObject, error) {
	return ec.stub("field-string") // Source file: editfns.c
}

func (ec *execContext) fieldStringNoProperties_autogen(pos lispObject) (lispObject, error) {
	return ec.stub("field-string-no-properties") // Source file: editfns.c
}

func (ec *execContext) followingChar_autogen() (lispObject, error) {
	return ec.stub("following-char") // Source file: editfns.c
}

func (ec *execContext) format_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("format") // Source file: editfns.c
}

func (ec *execContext) formatMessage_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("format-message") // Source file: editfns.c
}

func (ec *execContext) gapPosition_autogen() (lispObject, error) {
	return ec.stub("gap-position") // Source file: editfns.c
}

func (ec *execContext) gapSize_autogen() (lispObject, error) {
	return ec.stub("gap-size") // Source file: editfns.c
}

func (ec *execContext) getPosProperty_autogen(position, prop, object lispObject) (lispObject, error) {
	return ec.stub("get-pos-property") // Source file: editfns.c
}

func (ec *execContext) gotoChar_autogen(position lispObject) (lispObject, error) {
	return ec.stub("goto-char") // Source file: editfns.c
}

func (ec *execContext) groupGid_autogen() (lispObject, error) {
	return ec.stub("group-gid") // Source file: editfns.c
}

func (ec *execContext) groupName_autogen(gid lispObject) (lispObject, error) {
	return ec.stub("group-name") // Source file: editfns.c
}

func (ec *execContext) groupRealGid_autogen() (lispObject, error) {
	return ec.stub("group-real-gid") // Source file: editfns.c
}

func (ec *execContext) insert_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("insert") // Source file: editfns.c
}

func (ec *execContext) insertAndInherit_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("insert-and-inherit") // Source file: editfns.c
}

func (ec *execContext) insertBeforeMarkers_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("insert-before-markers") // Source file: editfns.c
}

func (ec *execContext) insertAndInheritBeforeMarkers_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("insert-before-markers-and-inherit") // Source file: editfns.c
}

func (ec *execContext) insertBufferSubstring_autogen(buffer, start, end lispObject) (lispObject, error) {
	return ec.stub("insert-buffer-substring") // Source file: editfns.c
}

func (ec *execContext) insertByte_autogen(byte, count, inherit lispObject) (lispObject, error) {
	return ec.stub("insert-byte") // Source file: editfns.c
}

func (ec *execContext) insertChar_autogen(character, count, inherit lispObject) (lispObject, error) {
	return ec.stub("insert-char") // Source file: editfns.c
}

func (ec *execContext) internalLockNarrowing_autogen(tag lispObject) (lispObject, error) {
	return ec.stub("internal--lock-narrowing") // Source file: editfns.c
}

func (ec *execContext) internalUnlockNarrowing_autogen(tag lispObject) (lispObject, error) {
	return ec.stub("internal--unlock-narrowing") // Source file: editfns.c
}

func (ec *execContext) lineBeginningPosition_autogen(n lispObject) (lispObject, error) {
	return ec.stub("line-beginning-position") // Source file: editfns.c
}

func (ec *execContext) lineEndPosition_autogen(n lispObject) (lispObject, error) {
	return ec.stub("line-end-position") // Source file: editfns.c
}

func (ec *execContext) markMarker_autogen() (lispObject, error) {
	return ec.stub("mark-marker") // Source file: editfns.c
}

func (ec *execContext) message_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("message") // Source file: editfns.c
}

func (ec *execContext) messageBox_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("message-box") // Source file: editfns.c
}

func (ec *execContext) messageOrBox_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("message-or-box") // Source file: editfns.c
}

func (ec *execContext) narrowToRegion_autogen(start, end lispObject) (lispObject, error) {
	return ec.stub("narrow-to-region") // Source file: editfns.c
}

func (ec *execContext) ngettext_autogen(msgid, msgid_plural, n lispObject) (lispObject, error) {
	return ec.stub("ngettext") // Source file: editfns.c
}

func (ec *execContext) point_autogen() (lispObject, error) {
	return ec.stub("point") // Source file: editfns.c
}

func (ec *execContext) pointMarker_autogen() (lispObject, error) {
	return ec.stub("point-marker") // Source file: editfns.c
}

func (ec *execContext) pointMax_autogen() (lispObject, error) {
	return ec.stub("point-max") // Source file: editfns.c
}

func (ec *execContext) pointMaxMarker_autogen() (lispObject, error) {
	return ec.stub("point-max-marker") // Source file: editfns.c
}

func (ec *execContext) pointMin_autogen() (lispObject, error) {
	return ec.stub("point-min") // Source file: editfns.c
}

func (ec *execContext) pointMinMarker_autogen() (lispObject, error) {
	return ec.stub("point-min-marker") // Source file: editfns.c
}

func (ec *execContext) posBol_autogen(n lispObject) (lispObject, error) {
	return ec.stub("pos-bol") // Source file: editfns.c
}

func (ec *execContext) posEol_autogen(n lispObject) (lispObject, error) {
	return ec.stub("pos-eol") // Source file: editfns.c
}

func (ec *execContext) positionBytes_autogen(position lispObject) (lispObject, error) {
	return ec.stub("position-bytes") // Source file: editfns.c
}

func (ec *execContext) previousChar_autogen() (lispObject, error) {
	return ec.stub("preceding-char") // Source file: editfns.c
}

func (ec *execContext) propertize_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("propertize") // Source file: editfns.c
}

func (ec *execContext) regionBeginning_autogen() (lispObject, error) {
	return ec.stub("region-beginning") // Source file: editfns.c
}

func (ec *execContext) regionEnd_autogen() (lispObject, error) {
	return ec.stub("region-end") // Source file: editfns.c
}

func (ec *execContext) replaceBufferContents_autogen(source, max_secs, max_costs lispObject) (lispObject, error) {
	return ec.stub("replace-buffer-contents") // Source file: editfns.c
}

func (ec *execContext) saveCurrentBuffer_autogen(args lispObject) (lispObject, error) {
	return ec.stub("save-current-buffer") // Source file: editfns.c
}

func (ec *execContext) saveExcursion_autogen(args lispObject) (lispObject, error) {
	return ec.stub("save-excursion") // Source file: editfns.c
}

func (ec *execContext) saveRestriction_autogen(body lispObject) (lispObject, error) {
	return ec.stub("save-restriction") // Source file: editfns.c
}

func (ec *execContext) stringToChar_autogen(string lispObject) (lispObject, error) {
	return ec.stub("string-to-char") // Source file: editfns.c
}

func (ec *execContext) substCharInRegion_autogen(start, end, fromchar, tochar, noundo lispObject) (lispObject, error) {
	return ec.stub("subst-char-in-region") // Source file: editfns.c
}

func (ec *execContext) systemName_autogen() (lispObject, error) {
	return ec.stub("system-name") // Source file: editfns.c
}

func (ec *execContext) translateRegionInternal_autogen(start, end, table lispObject) (lispObject, error) {
	return ec.stub("translate-region-internal") // Source file: editfns.c
}

func (ec *execContext) transposeRegions_autogen(startr1, endr1, startr2, endr2, leave_markers lispObject) (lispObject, error) {
	return ec.stub("transpose-regions") // Source file: editfns.c
}

func (ec *execContext) userFullName_autogen(uid lispObject) (lispObject, error) {
	return ec.stub("user-full-name") // Source file: editfns.c
}

func (ec *execContext) userLoginName_autogen(uid lispObject) (lispObject, error) {
	return ec.stub("user-login-name") // Source file: editfns.c
}

func (ec *execContext) userRealLoginName_autogen() (lispObject, error) {
	return ec.stub("user-real-login-name") // Source file: editfns.c
}

func (ec *execContext) userRealUid_autogen() (lispObject, error) {
	return ec.stub("user-real-uid") // Source file: editfns.c
}

func (ec *execContext) userUid_autogen() (lispObject, error) {
	return ec.stub("user-uid") // Source file: editfns.c
}

func (ec *execContext) widen_autogen() (lispObject, error) {
	return ec.stub("widen") // Source file: editfns.c
}

func (ec *execContext) moduleLoad_autogen(file lispObject) (lispObject, error) {
	return ec.stub("module-load") // Source file: emacs-module.c
}

func (ec *execContext) daemonInitialized_autogen() (lispObject, error) {
	return ec.stub("daemon-initialized") // Source file: emacs.c
}

func (ec *execContext) daemonp_autogen() (lispObject, error) {
	return ec.stub("daemonp") // Source file: emacs.c
}

func (ec *execContext) dumpEmacs_autogen(filename, symfile lispObject) (lispObject, error) {
	return ec.stub("dump-emacs") // Source file: emacs.c
}

func (ec *execContext) invocationDirectory_autogen() (lispObject, error) {
	return ec.stub("invocation-directory") // Source file: emacs.c
}

func (ec *execContext) invocationName_autogen() (lispObject, error) {
	return ec.stub("invocation-name") // Source file: emacs.c
}

func (ec *execContext) killEmacs_autogen(arg, restart lispObject) (lispObject, error) {
	return ec.stub("kill-emacs") // Source file: emacs.c, attributes: noreturn
}

func (ec *execContext) and_autogen(args lispObject) (lispObject, error) {
	return ec.stub("and") // Source file: eval.c
}

func (ec *execContext) apply_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("apply") // Source file: eval.c
}

func (ec *execContext) autoload_autogen(function, file, docstring, interactive, type_ lispObject) (lispObject, error) {
	return ec.stub("autoload") // Source file: eval.c
}

func (ec *execContext) autoloadDoLoad_autogen(fundef, funname, macro_only lispObject) (lispObject, error) {
	return ec.stub("autoload-do-load") // Source file: eval.c
}

func (ec *execContext) backtraceFramesFromThread_autogen(thread lispObject) (lispObject, error) {
	return ec.stub("backtrace--frames-from-thread") // Source file: eval.c
}

func (ec *execContext) backtraceLocals_autogen(nframes, base lispObject) (lispObject, error) {
	return ec.stub("backtrace--locals") // Source file: eval.c
}

func (ec *execContext) backtraceDebug_autogen(level, flag lispObject) (lispObject, error) {
	return ec.stub("backtrace-debug") // Source file: eval.c
}

func (ec *execContext) backtraceEval_autogen(exp, nframes, base lispObject) (lispObject, error) {
	return ec.stub("backtrace-eval") // Source file: eval.c
}

func (ec *execContext) backtraceFrameInternal_autogen(function, nframes, base lispObject) (lispObject, error) {
	return ec.stub("backtrace-frame--internal") // Source file: eval.c
}

func (ec *execContext) catch_autogen(args lispObject) (lispObject, error) {
	return ec.stub("catch") // Source file: eval.c
}

func (ec *execContext) commandp_autogen(function, for_call_interactively lispObject) (lispObject, error) {
	return ec.stub("commandp") // Source file: eval.c
}

func (ec *execContext) cond_autogen(args lispObject) (lispObject, error) {
	return ec.stub("cond") // Source file: eval.c
}

func (ec *execContext) conditionCase_autogen(args lispObject) (lispObject, error) {
	return ec.stub("condition-case") // Source file: eval.c
}

func (ec *execContext) defaultToplevelValue_autogen(symbol lispObject) (lispObject, error) {
	return ec.stub("default-toplevel-value") // Source file: eval.c
}

func (ec *execContext) defconst_autogen(args lispObject) (lispObject, error) {
	return ec.stub("defconst") // Source file: eval.c
}

func (ec *execContext) defconst1_autogen(sym, initvalue, docstring lispObject) (lispObject, error) {
	return ec.stub("defconst-1") // Source file: eval.c
}

func (ec *execContext) defvar_autogen(args lispObject) (lispObject, error) {
	return ec.stub("defvar") // Source file: eval.c
}

func (ec *execContext) defvar1_autogen(sym, initvalue, docstring lispObject) (lispObject, error) {
	return ec.stub("defvar-1") // Source file: eval.c
}

func (ec *execContext) defvaralias_autogen(new_alias, base_variable, docstring lispObject) (lispObject, error) {
	return ec.stub("defvaralias") // Source file: eval.c
}

func (ec *execContext) eval_autogen(form, lexical lispObject) (lispObject, error) {
	return ec.stub("eval") // Source file: eval.c
}

func (ec *execContext) fetchBytecode_autogen(object lispObject) (lispObject, error) {
	return ec.stub("fetch-bytecode") // Source file: eval.c
}

func (ec *execContext) funcArity_autogen(function lispObject) (lispObject, error) {
	return ec.stub("func-arity") // Source file: eval.c
}

func (ec *execContext) funcall_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("funcall") // Source file: eval.c
}

func (ec *execContext) funcallWithDelayedMessage_autogen(timeout, message, function lispObject) (lispObject, error) {
	return ec.stub("funcall-with-delayed-message") // Source file: eval.c
}

func (ec *execContext) function_autogen(args lispObject) (lispObject, error) {
	return ec.stub("function") // Source file: eval.c
}

func (ec *execContext) functionp_autogen(object lispObject) (lispObject, error) {
	return ec.stub("functionp") // Source file: eval.c
}

func (ec *execContext) if_autogen(args lispObject) (lispObject, error) {
	return ec.stub("if") // Source file: eval.c
}

func (ec *execContext) internalDefineUninitializedVariable_autogen(symbol, doc lispObject) (lispObject, error) {
	return ec.stub("internal--define-uninitialized-variable") // Source file: eval.c
}

func (ec *execContext) makeVarNonSpecial_autogen(symbol lispObject) (lispObject, error) {
	return ec.stub("internal-make-var-non-special") // Source file: eval.c
}

func (ec *execContext) let_autogen(args lispObject) (lispObject, error) {
	return ec.stub("let") // Source file: eval.c
}

func (ec *execContext) letX_autogen(args lispObject) (lispObject, error) {
	return ec.stub("let*") // Source file: eval.c
}

func (ec *execContext) macroexpand_autogen(form, environment lispObject) (lispObject, error) {
	return ec.stub("macroexpand") // Source file: eval.c
}

func (ec *execContext) mapbacktrace_autogen(function, base lispObject) (lispObject, error) {
	return ec.stub("mapbacktrace") // Source file: eval.c
}

func (ec *execContext) or_autogen(args lispObject) (lispObject, error) {
	return ec.stub("or") // Source file: eval.c
}

func (ec *execContext) prog1_autogen(args lispObject) (lispObject, error) {
	return ec.stub("prog1") // Source file: eval.c
}

func (ec *execContext) progn_autogen(body lispObject) (lispObject, error) {
	return ec.stub("progn") // Source file: eval.c
}

func (ec *execContext) quote_autogen(args lispObject) (lispObject, error) {
	return ec.stub("quote") // Source file: eval.c
}

func (ec *execContext) runHookWithArgs_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("run-hook-with-args") // Source file: eval.c
}

func (ec *execContext) runHookWithArgsUntilFailure_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("run-hook-with-args-until-failure") // Source file: eval.c
}

func (ec *execContext) runHookWithArgsUntilSuccess_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("run-hook-with-args-until-success") // Source file: eval.c
}

func (ec *execContext) runHookWrapped_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("run-hook-wrapped") // Source file: eval.c
}

func (ec *execContext) runHooks_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("run-hooks") // Source file: eval.c
}

func (ec *execContext) setDefaultToplevelValue_autogen(symbol, value lispObject) (lispObject, error) {
	return ec.stub("set-default-toplevel-value") // Source file: eval.c
}

func (ec *execContext) setq_autogen(args lispObject) (lispObject, error) {
	return ec.stub("setq") // Source file: eval.c
}

func (ec *execContext) signal_autogen(error_symbol, data lispObject) (lispObject, error) {
	return ec.stub("signal") // Source file: eval.c, attributes: noreturn
}

func (ec *execContext) specialVariableP_autogen(symbol lispObject) (lispObject, error) {
	return ec.stub("special-variable-p") // Source file: eval.c
}

func (ec *execContext) throw_autogen(tag, value lispObject) (lispObject, error) {
	return ec.stub("throw") // Source file: eval.c, attributes: noreturn
}

func (ec *execContext) unwindProtect_autogen(args lispObject) (lispObject, error) {
	return ec.stub("unwind-protect") // Source file: eval.c
}

func (ec *execContext) while_autogen(args lispObject) (lispObject, error) {
	return ec.stub("while") // Source file: eval.c
}

func (ec *execContext) accessFile_autogen(filename, string lispObject) (lispObject, error) {
	return ec.stub("access-file") // Source file: fileio.c
}

func (ec *execContext) addNameToFile_autogen(file, newname, ok_if_already_exists lispObject) (lispObject, error) {
	return ec.stub("add-name-to-file") // Source file: fileio.c
}

func (ec *execContext) carLessThanCar_autogen(a, b lispObject) (lispObject, error) {
	return ec.stub("car-less-than-car") // Source file: fileio.c
}

func (ec *execContext) clearBufferAutoSaveFailure_autogen() (lispObject, error) {
	return ec.stub("clear-buffer-auto-save-failure") // Source file: fileio.c
}

func (ec *execContext) copyFile_autogen(file, newname, ok_if_already_exists, keep_time, preserve_uid_gid, preserve_permissions lispObject) (lispObject, error) {
	return ec.stub("copy-file") // Source file: fileio.c
}

func (ec *execContext) defaultFileModes_autogen() (lispObject, error) {
	return ec.stub("default-file-modes") // Source file: fileio.c
}

func (ec *execContext) deleteDirectoryInternal_autogen(directory lispObject) (lispObject, error) {
	return ec.stub("delete-directory-internal") // Source file: fileio.c
}

func (ec *execContext) deleteFile_autogen(filename, trash lispObject) (lispObject, error) {
	return ec.stub("delete-file") // Source file: fileio.c
}

func (ec *execContext) directoryFileName_autogen(directory lispObject) (lispObject, error) {
	return ec.stub("directory-file-name") // Source file: fileio.c
}

func (ec *execContext) directoryNameP_autogen(name lispObject) (lispObject, error) {
	return ec.stub("directory-name-p") // Source file: fileio.c
}

func (ec *execContext) doAutoSave_autogen(no_message, current_only lispObject) (lispObject, error) {
	return ec.stub("do-auto-save") // Source file: fileio.c
}

func (ec *execContext) expandFileName_autogen(name, default_directory lispObject) (lispObject, error) {
	return ec.stub("expand-file-name") // Source file: fileio.c
}

func (ec *execContext) fileAccessibleDirectoryP_autogen(filename lispObject) (lispObject, error) {
	return ec.stub("file-accessible-directory-p") // Source file: fileio.c
}

func (ec *execContext) fileAcl_autogen(filename lispObject) (lispObject, error) {
	return ec.stub("file-acl") // Source file: fileio.c
}

func (ec *execContext) fileDirectoryP_autogen(filename lispObject) (lispObject, error) {
	return ec.stub("file-directory-p") // Source file: fileio.c
}

func (ec *execContext) fileExecutableP_autogen(filename lispObject) (lispObject, error) {
	return ec.stub("file-executable-p") // Source file: fileio.c
}

func (ec *execContext) fileExistsP_autogen(filename lispObject) (lispObject, error) {
	return ec.stub("file-exists-p") // Source file: fileio.c
}

func (ec *execContext) fileModes_autogen(filename, flag lispObject) (lispObject, error) {
	return ec.stub("file-modes") // Source file: fileio.c
}

func (ec *execContext) fileNameAbsoluteP_autogen(filename lispObject) (lispObject, error) {
	return ec.stub("file-name-absolute-p") // Source file: fileio.c
}

func (ec *execContext) fileNameAsDirectory_autogen(file lispObject) (lispObject, error) {
	return ec.stub("file-name-as-directory") // Source file: fileio.c
}

func (ec *execContext) fileNameCaseInsensitiveP_autogen(filename lispObject) (lispObject, error) {
	return ec.stub("file-name-case-insensitive-p") // Source file: fileio.c
}

func (ec *execContext) fileNameConcat_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("file-name-concat") // Source file: fileio.c
}

func (ec *execContext) fileNameDirectory_autogen(filename lispObject) (lispObject, error) {
	return ec.stub("file-name-directory") // Source file: fileio.c
}

func (ec *execContext) fileNameNondirectory_autogen(filename lispObject) (lispObject, error) {
	return ec.stub("file-name-nondirectory") // Source file: fileio.c
}

func (ec *execContext) fileNewerThanFileP_autogen(file1, file2 lispObject) (lispObject, error) {
	return ec.stub("file-newer-than-file-p") // Source file: fileio.c
}

func (ec *execContext) fileReadableP_autogen(filename lispObject) (lispObject, error) {
	return ec.stub("file-readable-p") // Source file: fileio.c
}

func (ec *execContext) fileRegularP_autogen(filename lispObject) (lispObject, error) {
	return ec.stub("file-regular-p") // Source file: fileio.c
}

func (ec *execContext) fileSelinuxContext_autogen(filename lispObject) (lispObject, error) {
	return ec.stub("file-selinux-context") // Source file: fileio.c
}

func (ec *execContext) fileSymlinkP_autogen(filename lispObject) (lispObject, error) {
	return ec.stub("file-symlink-p") // Source file: fileio.c
}

func (ec *execContext) fileWritableP_autogen(filename lispObject) (lispObject, error) {
	return ec.stub("file-writable-p") // Source file: fileio.c
}

func (ec *execContext) findFileNameHandler_autogen(filename, operation lispObject) (lispObject, error) {
	return ec.stub("find-file-name-handler") // Source file: fileio.c
}

func (ec *execContext) insertFileContents_autogen(filename, visit, beg, end, replace lispObject) (lispObject, error) {
	return ec.stub("insert-file-contents") // Source file: fileio.c
}

func (ec *execContext) makeDirectoryInternal_autogen(directory lispObject) (lispObject, error) {
	return ec.stub("make-directory-internal") // Source file: fileio.c
}

func (ec *execContext) makeSymbolicLink_autogen(target, linkname, ok_if_already_exists lispObject) (lispObject, error) {
	return ec.stub("make-symbolic-link") // Source file: fileio.c
}

func (ec *execContext) makeTempFileInternal_autogen(prefix, dir_flag, suffix, text lispObject) (lispObject, error) {
	return ec.stub("make-temp-file-internal") // Source file: fileio.c
}

func (ec *execContext) makeTempName_autogen(prefix lispObject) (lispObject, error) {
	return ec.stub("make-temp-name") // Source file: fileio.c
}

func (ec *execContext) nextReadFileUsesDialogP_autogen() (lispObject, error) {
	return ec.stub("next-read-file-uses-dialog-p") // Source file: fileio.c
}

func (ec *execContext) recentAutoSaveP_autogen() (lispObject, error) {
	return ec.stub("recent-auto-save-p") // Source file: fileio.c
}

func (ec *execContext) renameFile_autogen(file, newname, ok_if_already_exists lispObject) (lispObject, error) {
	return ec.stub("rename-file") // Source file: fileio.c
}

func (ec *execContext) setBinaryMode_autogen(stream, mode lispObject) (lispObject, error) {
	return ec.stub("set-binary-mode") // Source file: fileio.c
}

func (ec *execContext) setBufferAutoSaved_autogen() (lispObject, error) {
	return ec.stub("set-buffer-auto-saved") // Source file: fileio.c
}

func (ec *execContext) setDefaultFileModes_autogen(mode lispObject) (lispObject, error) {
	return ec.stub("set-default-file-modes") // Source file: fileio.c
}

func (ec *execContext) setFileAcl_autogen(filename, acl_string lispObject) (lispObject, error) {
	return ec.stub("set-file-acl") // Source file: fileio.c
}

func (ec *execContext) setFileModes_autogen(filename, mode, flag lispObject) (lispObject, error) {
	return ec.stub("set-file-modes") // Source file: fileio.c
}

func (ec *execContext) setFileSelinuxContext_autogen(filename, context lispObject) (lispObject, error) {
	return ec.stub("set-file-selinux-context") // Source file: fileio.c
}

func (ec *execContext) setFileTimes_autogen(filename, timestamp, flag lispObject) (lispObject, error) {
	return ec.stub("set-file-times") // Source file: fileio.c
}

func (ec *execContext) setVisitedFileModtime_autogen(time_flag lispObject) (lispObject, error) {
	return ec.stub("set-visited-file-modtime") // Source file: fileio.c
}

func (ec *execContext) substituteInFileName_autogen(filename lispObject) (lispObject, error) {
	return ec.stub("substitute-in-file-name") // Source file: fileio.c
}

func (ec *execContext) unhandledFileNameDirectory_autogen(filename lispObject) (lispObject, error) {
	return ec.stub("unhandled-file-name-directory") // Source file: fileio.c
}

func (ec *execContext) unixSync_autogen() (lispObject, error) {
	return ec.stub("unix-sync") // Source file: fileio.c
}

func (ec *execContext) verifyVisitedFileModtime_autogen(buf lispObject) (lispObject, error) {
	return ec.stub("verify-visited-file-modtime") // Source file: fileio.c
}

func (ec *execContext) visitedFileModtime_autogen() (lispObject, error) {
	return ec.stub("visited-file-modtime") // Source file: fileio.c
}

func (ec *execContext) writeRegion_autogen(start, end, filename, append, visit, lockname, mustbenew lispObject) (lispObject, error) {
	return ec.stub("write-region") // Source file: fileio.c
}

func (ec *execContext) fileLockedP_autogen(filename lispObject) (lispObject, error) {
	return ec.stub("file-locked-p") // Source file: filelock.c
}

func (ec *execContext) lockBuffer_autogen(file lispObject) (lispObject, error) {
	return ec.stub("lock-buffer") // Source file: filelock.c
}

func (ec *execContext) lockFile_autogen(file lispObject) (lispObject, error) {
	return ec.stub("lock-file") // Source file: filelock.c
}

func (ec *execContext) unlockBuffer_autogen() (lispObject, error) {
	return ec.stub("unlock-buffer") // Source file: filelock.c
}

func (ec *execContext) unlockFile_autogen(file lispObject) (lispObject, error) {
	return ec.stub("unlock-file") // Source file: filelock.c
}

func (ec *execContext) abs_autogen(arg lispObject) (lispObject, error) {
	return ec.stub("abs") // Source file: floatfns.c
}

func (ec *execContext) acos_autogen(arg lispObject) (lispObject, error) {
	return ec.stub("acos") // Source file: floatfns.c
}

func (ec *execContext) asin_autogen(arg lispObject) (lispObject, error) {
	return ec.stub("asin") // Source file: floatfns.c
}

func (ec *execContext) atan_autogen(y, x lispObject) (lispObject, error) {
	return ec.stub("atan") // Source file: floatfns.c
}

func (ec *execContext) ceiling_autogen(arg, divisor lispObject) (lispObject, error) {
	return ec.stub("ceiling") // Source file: floatfns.c
}

func (ec *execContext) copysign_autogen(x1, x2 lispObject) (lispObject, error) {
	return ec.stub("copysign") // Source file: floatfns.c
}

func (ec *execContext) cos_autogen(arg lispObject) (lispObject, error) {
	return ec.stub("cos") // Source file: floatfns.c
}

func (ec *execContext) exp_autogen(arg lispObject) (lispObject, error) {
	return ec.stub("exp") // Source file: floatfns.c
}

func (ec *execContext) expt_autogen(arg1, arg2 lispObject) (lispObject, error) {
	return ec.stub("expt") // Source file: floatfns.c
}

func (ec *execContext) fceiling_autogen(arg lispObject) (lispObject, error) {
	return ec.stub("fceiling") // Source file: floatfns.c
}

func (ec *execContext) ffloor_autogen(arg lispObject) (lispObject, error) {
	return ec.stub("ffloor") // Source file: floatfns.c
}

func (ec *execContext) float_autogen(arg lispObject) (lispObject, error) {
	return ec.stub("float") // Source file: floatfns.c
}

func (ec *execContext) floor_autogen(arg, divisor lispObject) (lispObject, error) {
	return ec.stub("floor") // Source file: floatfns.c
}

func (ec *execContext) frexp_autogen(x lispObject) (lispObject, error) {
	return ec.stub("frexp") // Source file: floatfns.c
}

func (ec *execContext) fround_autogen(arg lispObject) (lispObject, error) {
	return ec.stub("fround") // Source file: floatfns.c
}

func (ec *execContext) ftruncate_autogen(arg lispObject) (lispObject, error) {
	return ec.stub("ftruncate") // Source file: floatfns.c
}

func (ec *execContext) isnan_autogen(x lispObject) (lispObject, error) {
	return ec.stub("isnan") // Source file: floatfns.c
}

func (ec *execContext) ldexp_autogen(sgnfcand, exponent lispObject) (lispObject, error) {
	return ec.stub("ldexp") // Source file: floatfns.c
}

func (ec *execContext) log_autogen(arg, base lispObject) (lispObject, error) {
	return ec.stub("log") // Source file: floatfns.c
}

func (ec *execContext) logb_autogen(arg lispObject) (lispObject, error) {
	return ec.stub("logb") // Source file: floatfns.c
}

func (ec *execContext) round_autogen(arg, divisor lispObject) (lispObject, error) {
	return ec.stub("round") // Source file: floatfns.c
}

func (ec *execContext) sin_autogen(arg lispObject) (lispObject, error) {
	return ec.stub("sin") // Source file: floatfns.c
}

func (ec *execContext) sqrt_autogen(arg lispObject) (lispObject, error) {
	return ec.stub("sqrt") // Source file: floatfns.c
}

func (ec *execContext) tan_autogen(arg lispObject) (lispObject, error) {
	return ec.stub("tan") // Source file: floatfns.c
}

func (ec *execContext) truncate_autogen(arg, divisor lispObject) (lispObject, error) {
	return ec.stub("truncate") // Source file: floatfns.c
}

func (ec *execContext) append_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("append") // Source file: fns.c
}

func (ec *execContext) assoc_autogen(key, alist, testfn lispObject) (lispObject, error) {
	return ec.stub("assoc") // Source file: fns.c
}

func (ec *execContext) assq_autogen(key, alist lispObject) (lispObject, error) {
	return ec.stub("assq") // Source file: fns.c
}

func (ec *execContext) base64DecodeRegion_autogen(beg, end, base64url, ignore_invalid lispObject) (lispObject, error) {
	return ec.stub("base64-decode-region") // Source file: fns.c
}

func (ec *execContext) base64DecodeString_autogen(string, base64url, ignore_invalid lispObject) (lispObject, error) {
	return ec.stub("base64-decode-string") // Source file: fns.c
}

func (ec *execContext) base64EncodeRegion_autogen(beg, end, no_line_break lispObject) (lispObject, error) {
	return ec.stub("base64-encode-region") // Source file: fns.c
}

func (ec *execContext) base64EncodeString_autogen(string, no_line_break lispObject) (lispObject, error) {
	return ec.stub("base64-encode-string") // Source file: fns.c
}

func (ec *execContext) base64urlEncodeRegion_autogen(beg, end, no_pad lispObject) (lispObject, error) {
	return ec.stub("base64url-encode-region") // Source file: fns.c
}

func (ec *execContext) base64urlEncodeString_autogen(string, no_pad lispObject) (lispObject, error) {
	return ec.stub("base64url-encode-string") // Source file: fns.c
}

func (ec *execContext) bufferHash_autogen(buffer_or_name lispObject) (lispObject, error) {
	return ec.stub("buffer-hash") // Source file: fns.c
}

func (ec *execContext) bufferLineStatistics_autogen(buffer_or_name lispObject) (lispObject, error) {
	return ec.stub("buffer-line-statistics") // Source file: fns.c
}

func (ec *execContext) clearString_autogen(string lispObject) (lispObject, error) {
	return ec.stub("clear-string") // Source file: fns.c
}

func (ec *execContext) clrhash_autogen(table lispObject) (lispObject, error) {
	return ec.stub("clrhash") // Source file: fns.c
}

func (ec *execContext) compareStrings_autogen(str1, start1, end1, str2, start2, end2, ignore_case lispObject) (lispObject, error) {
	return ec.stub("compare-strings") // Source file: fns.c
}

func (ec *execContext) concat_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("concat") // Source file: fns.c
}

func (ec *execContext) copyAlist_autogen(alist lispObject) (lispObject, error) {
	return ec.stub("copy-alist") // Source file: fns.c
}

func (ec *execContext) copyHashTable_autogen(table lispObject) (lispObject, error) {
	return ec.stub("copy-hash-table") // Source file: fns.c
}

func (ec *execContext) copySequence_autogen(arg lispObject) (lispObject, error) {
	return ec.stub("copy-sequence") // Source file: fns.c
}

func (ec *execContext) defineHashTableTest_autogen(name, test, hash lispObject) (lispObject, error) {
	return ec.stub("define-hash-table-test") // Source file: fns.c
}

func (ec *execContext) delete_autogen(elt, seq lispObject) (lispObject, error) {
	return ec.stub("delete") // Source file: fns.c
}

func (ec *execContext) delq_autogen(elt, list lispObject) (lispObject, error) {
	return ec.stub("delq") // Source file: fns.c
}

func (ec *execContext) elt_autogen(sequence, n lispObject) (lispObject, error) {
	return ec.stub("elt") // Source file: fns.c
}

func (ec *execContext) eql_autogen(obj1, obj2 lispObject) (lispObject, error) {
	return ec.stub("eql") // Source file: fns.c
}

func (ec *execContext) equal_autogen(o1, o2 lispObject) (lispObject, error) {
	return ec.stub("equal") // Source file: fns.c
}

func (ec *execContext) equalIncludingProperties_autogen(o1, o2 lispObject) (lispObject, error) {
	return ec.stub("equal-including-properties") // Source file: fns.c
}

func (ec *execContext) featurep_autogen(feature, subfeature lispObject) (lispObject, error) {
	return ec.stub("featurep") // Source file: fns.c
}

func (ec *execContext) fillarray_autogen(array, item lispObject) (lispObject, error) {
	return ec.stub("fillarray") // Source file: fns.c
}

func (ec *execContext) get_autogen(symbol, propname lispObject) (lispObject, error) {
	return ec.stub("get") // Source file: fns.c
}

func (ec *execContext) gethash_autogen(key, table, dflt lispObject) (lispObject, error) {
	return ec.stub("gethash") // Source file: fns.c
}

func (ec *execContext) hashTableCount_autogen(table lispObject) (lispObject, error) {
	return ec.stub("hash-table-count") // Source file: fns.c
}

func (ec *execContext) hashTableP_autogen(obj lispObject) (lispObject, error) {
	return ec.stub("hash-table-p") // Source file: fns.c
}

func (ec *execContext) hashTableRehashSize_autogen(table lispObject) (lispObject, error) {
	return ec.stub("hash-table-rehash-size") // Source file: fns.c
}

func (ec *execContext) hashTableRehashThreshold_autogen(table lispObject) (lispObject, error) {
	return ec.stub("hash-table-rehash-threshold") // Source file: fns.c
}

func (ec *execContext) hashTableSize_autogen(table lispObject) (lispObject, error) {
	return ec.stub("hash-table-size") // Source file: fns.c
}

func (ec *execContext) hashTableTest_autogen(table lispObject) (lispObject, error) {
	return ec.stub("hash-table-test") // Source file: fns.c
}

func (ec *execContext) hashTableWeakness_autogen(table lispObject) (lispObject, error) {
	return ec.stub("hash-table-weakness") // Source file: fns.c
}

func (ec *execContext) identity_autogen(argument lispObject) (lispObject, error) {
	return ec.stub("identity") // Source file: fns.c, attributes: const
}

func (ec *execContext) length_autogen(sequence lispObject) (lispObject, error) {
	return ec.stub("length") // Source file: fns.c
}

func (ec *execContext) lengthLess_autogen(sequence, length lispObject) (lispObject, error) {
	return ec.stub("length<") // Source file: fns.c
}

func (ec *execContext) lengthEqual_autogen(sequence, length lispObject) (lispObject, error) {
	return ec.stub("length=") // Source file: fns.c
}

func (ec *execContext) lengthGreater_autogen(sequence, length lispObject) (lispObject, error) {
	return ec.stub("length>") // Source file: fns.c
}

func (ec *execContext) lineNumberAtPos_autogen(position, absolute lispObject) (lispObject, error) {
	return ec.stub("line-number-at-pos") // Source file: fns.c
}

func (ec *execContext) loadAverage_autogen(use_floats lispObject) (lispObject, error) {
	return ec.stub("load-average") // Source file: fns.c
}

func (ec *execContext) localeInfo_autogen(item lispObject) (lispObject, error) {
	return ec.stub("locale-info") // Source file: fns.c
}

func (ec *execContext) makeHashTable_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("make-hash-table") // Source file: fns.c
}

func (ec *execContext) mapc_autogen(function, sequence lispObject) (lispObject, error) {
	return ec.stub("mapc") // Source file: fns.c
}

func (ec *execContext) mapcan_autogen(function, sequence lispObject) (lispObject, error) {
	return ec.stub("mapcan") // Source file: fns.c
}

func (ec *execContext) mapcar_autogen(function, sequence lispObject) (lispObject, error) {
	return ec.stub("mapcar") // Source file: fns.c
}

func (ec *execContext) mapconcat_autogen(function, sequence, separator lispObject) (lispObject, error) {
	return ec.stub("mapconcat") // Source file: fns.c
}

func (ec *execContext) maphash_autogen(function, table lispObject) (lispObject, error) {
	return ec.stub("maphash") // Source file: fns.c
}

func (ec *execContext) md5_autogen(object, start, end, coding_system, noerror lispObject) (lispObject, error) {
	return ec.stub("md5") // Source file: fns.c
}

func (ec *execContext) member_autogen(elt, list lispObject) (lispObject, error) {
	return ec.stub("member") // Source file: fns.c
}

func (ec *execContext) memq_autogen(elt, list lispObject) (lispObject, error) {
	return ec.stub("memq") // Source file: fns.c
}

func (ec *execContext) memql_autogen(elt, list lispObject) (lispObject, error) {
	return ec.stub("memql") // Source file: fns.c
}

func (ec *execContext) nconc_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("nconc") // Source file: fns.c
}

func (ec *execContext) nreverse_autogen(seq lispObject) (lispObject, error) {
	return ec.stub("nreverse") // Source file: fns.c
}

func (ec *execContext) ntake_autogen(n, list lispObject) (lispObject, error) {
	return ec.stub("ntake") // Source file: fns.c
}

func (ec *execContext) nth_autogen(n, list lispObject) (lispObject, error) {
	return ec.stub("nth") // Source file: fns.c
}

func (ec *execContext) nthcdr_autogen(n, list lispObject) (lispObject, error) {
	return ec.stub("nthcdr") // Source file: fns.c
}

func (ec *execContext) objectIntervals_autogen(object lispObject) (lispObject, error) {
	return ec.stub("object-intervals") // Source file: fns.c
}

func (ec *execContext) plistGet_autogen(plist, prop, predicate lispObject) (lispObject, error) {
	return ec.stub("plist-get") // Source file: fns.c
}

func (ec *execContext) plistMember_autogen(plist, prop, predicate lispObject) (lispObject, error) {
	return ec.stub("plist-member") // Source file: fns.c
}

func (ec *execContext) plistPut_autogen(plist, prop, val, predicate lispObject) (lispObject, error) {
	return ec.stub("plist-put") // Source file: fns.c
}

func (ec *execContext) properListP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("proper-list-p") // Source file: fns.c, attributes: const
}

func (ec *execContext) provide_autogen(feature, subfeatures lispObject) (lispObject, error) {
	return ec.stub("provide") // Source file: fns.c
}

func (ec *execContext) put_autogen(symbol, propname, value lispObject) (lispObject, error) {
	return ec.stub("put") // Source file: fns.c
}

func (ec *execContext) puthash_autogen(key, value, table lispObject) (lispObject, error) {
	return ec.stub("puthash") // Source file: fns.c
}

func (ec *execContext) random_autogen(limit lispObject) (lispObject, error) {
	return ec.stub("random") // Source file: fns.c
}

func (ec *execContext) rassoc_autogen(key, alist lispObject) (lispObject, error) {
	return ec.stub("rassoc") // Source file: fns.c
}

func (ec *execContext) rassq_autogen(key, alist lispObject) (lispObject, error) {
	return ec.stub("rassq") // Source file: fns.c
}

func (ec *execContext) remhash_autogen(key, table lispObject) (lispObject, error) {
	return ec.stub("remhash") // Source file: fns.c
}

func (ec *execContext) require_autogen(feature, filename, noerror lispObject) (lispObject, error) {
	return ec.stub("require") // Source file: fns.c
}

func (ec *execContext) reverse_autogen(seq lispObject) (lispObject, error) {
	return ec.stub("reverse") // Source file: fns.c
}

func (ec *execContext) safeLength_autogen(list lispObject) (lispObject, error) {
	return ec.stub("safe-length") // Source file: fns.c
}

func (ec *execContext) secureHash_autogen(algorithm, object, start, end, binary lispObject) (lispObject, error) {
	return ec.stub("secure-hash") // Source file: fns.c
}

func (ec *execContext) secureHashAlgorithms_autogen() (lispObject, error) {
	return ec.stub("secure-hash-algorithms") // Source file: fns.c
}

func (ec *execContext) sort_autogen(seq, predicate lispObject) (lispObject, error) {
	return ec.stub("sort") // Source file: fns.c
}

func (ec *execContext) stringAsMultibyte_autogen(string lispObject) (lispObject, error) {
	return ec.stub("string-as-multibyte") // Source file: fns.c
}

func (ec *execContext) stringAsUnibyte_autogen(string lispObject) (lispObject, error) {
	return ec.stub("string-as-unibyte") // Source file: fns.c
}

func (ec *execContext) stringBytes_autogen(string lispObject) (lispObject, error) {
	return ec.stub("string-bytes") // Source file: fns.c
}

func (ec *execContext) stringCollateEqualp_autogen(s1, s2, locale, ignore_case lispObject) (lispObject, error) {
	return ec.stub("string-collate-equalp") // Source file: fns.c
}

func (ec *execContext) stringCollateLessp_autogen(s1, s2, locale, ignore_case lispObject) (lispObject, error) {
	return ec.stub("string-collate-lessp") // Source file: fns.c
}

func (ec *execContext) stringDistance_autogen(string1, string2, bytecompare lispObject) (lispObject, error) {
	return ec.stub("string-distance") // Source file: fns.c
}

func (ec *execContext) stringEqual_autogen(s1, s2 lispObject) (lispObject, error) {
	return ec.stub("string-equal") // Source file: fns.c
}

func (ec *execContext) stringLessp_autogen(string1, string2 lispObject) (lispObject, error) {
	return ec.stub("string-lessp") // Source file: fns.c
}

func (ec *execContext) stringMakeMultibyte_autogen(string lispObject) (lispObject, error) {
	return ec.stub("string-make-multibyte") // Source file: fns.c
}

func (ec *execContext) stringMakeUnibyte_autogen(string lispObject) (lispObject, error) {
	return ec.stub("string-make-unibyte") // Source file: fns.c
}

func (ec *execContext) stringSearch_autogen(needle, haystack, start_pos lispObject) (lispObject, error) {
	return ec.stub("string-search") // Source file: fns.c
}

func (ec *execContext) stringToMultibyte_autogen(string lispObject) (lispObject, error) {
	return ec.stub("string-to-multibyte") // Source file: fns.c
}

func (ec *execContext) stringToUnibyte_autogen(string lispObject) (lispObject, error) {
	return ec.stub("string-to-unibyte") // Source file: fns.c
}

func (ec *execContext) stringVersionLessp_autogen(string1, string2 lispObject) (lispObject, error) {
	return ec.stub("string-version-lessp") // Source file: fns.c
}

func (ec *execContext) substring_autogen(string, from, to lispObject) (lispObject, error) {
	return ec.stub("substring") // Source file: fns.c
}

func (ec *execContext) substringNoProperties_autogen(string, from, to lispObject) (lispObject, error) {
	return ec.stub("substring-no-properties") // Source file: fns.c
}

func (ec *execContext) sxhashEq_autogen(obj lispObject) (lispObject, error) {
	return ec.stub("sxhash-eq") // Source file: fns.c
}

func (ec *execContext) sxhashEql_autogen(obj lispObject) (lispObject, error) {
	return ec.stub("sxhash-eql") // Source file: fns.c
}

func (ec *execContext) sxhashEqual_autogen(obj lispObject) (lispObject, error) {
	return ec.stub("sxhash-equal") // Source file: fns.c
}

func (ec *execContext) sxhashEqualIncludingProperties_autogen(obj lispObject) (lispObject, error) {
	return ec.stub("sxhash-equal-including-properties") // Source file: fns.c
}

func (ec *execContext) take_autogen(n, list lispObject) (lispObject, error) {
	return ec.stub("take") // Source file: fns.c
}

func (ec *execContext) vconcat_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("vconcat") // Source file: fns.c
}

func (ec *execContext) widgetApply_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("widget-apply") // Source file: fns.c
}

func (ec *execContext) widgetGet_autogen(widget, property lispObject) (lispObject, error) {
	return ec.stub("widget-get") // Source file: fns.c
}

func (ec *execContext) widgetPut_autogen(widget, property, value lispObject) (lispObject, error) {
	return ec.stub("widget-put") // Source file: fns.c
}

func (ec *execContext) yesOrNoP_autogen(prompt lispObject) (lispObject, error) {
	return ec.stub("yes-or-no-p") // Source file: fns.c
}

func (ec *execContext) clearFontCache_autogen() (lispObject, error) {
	return ec.stub("clear-font-cache") // Source file: font.c
}

func (ec *execContext) closeFont_autogen(font_object, frame lispObject) (lispObject, error) {
	return ec.stub("close-font") // Source file: font.c
}

func (ec *execContext) drawString_autogen(font_object, string lispObject) (lispObject, error) {
	return ec.stub("draw-string") // Source file: font.c
}

func (ec *execContext) findFont_autogen(font_spec, frame lispObject) (lispObject, error) {
	return ec.stub("find-font") // Source file: font.c
}

func (ec *execContext) fontAt_autogen(position, window, string lispObject) (lispObject, error) {
	return ec.stub("font-at") // Source file: font.c
}

func (ec *execContext) fontDriveOtf_autogen(otf_features, gstring_in, from, to, gstring_out, index lispObject) (lispObject, error) {
	return ec.stub("font-drive-otf") // Source file: font.c
}

func (ec *execContext) fontFaceAttributes_autogen(font, frame lispObject) (lispObject, error) {
	return ec.stub("font-face-attributes") // Source file: font.c
}

func (ec *execContext) fontFamilyList_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("font-family-list") // Source file: font.c
}

func (ec *execContext) fontGet_autogen(font, key lispObject) (lispObject, error) {
	return ec.stub("font-get") // Source file: font.c
}

func (ec *execContext) fontGetGlyphs_autogen(font_object, from, to, object lispObject) (lispObject, error) {
	return ec.stub("font-get-glyphs") // Source file: font.c
}

func (ec *execContext) fontHasCharP_autogen(font, ch, frame lispObject) (lispObject, error) {
	return ec.stub("font-has-char-p") // Source file: font.c
}

func (ec *execContext) fontInfo_autogen(name, frame lispObject) (lispObject, error) {
	return ec.stub("font-info") // Source file: font.c
}

func (ec *execContext) fontMatchP_autogen(spec, font lispObject) (lispObject, error) {
	return ec.stub("font-match-p") // Source file: font.c
}

func (ec *execContext) fontOtfAlternates_autogen(font_object, character, otf_features lispObject) (lispObject, error) {
	return ec.stub("font-otf-alternates") // Source file: font.c
}

func (ec *execContext) fontPut_autogen(font, prop, val lispObject) (lispObject, error) {
	return ec.stub("font-put") // Source file: font.c
}

func (ec *execContext) fontShapeGstring_autogen(gstring, direction lispObject) (lispObject, error) {
	return ec.stub("font-shape-gstring") // Source file: font.c
}

func (ec *execContext) fontSpec_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("font-spec") // Source file: font.c
}

func (ec *execContext) fontVariationGlyphs_autogen(font_object, character lispObject) (lispObject, error) {
	return ec.stub("font-variation-glyphs") // Source file: font.c
}

func (ec *execContext) fontXlfdName_autogen(font, fold_wildcards lispObject) (lispObject, error) {
	return ec.stub("font-xlfd-name") // Source file: font.c
}

func (ec *execContext) fontp_autogen(object, extra_type lispObject) (lispObject, error) {
	return ec.stub("fontp") // Source file: font.c
}

func (ec *execContext) frameFontCache_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("frame-font-cache") // Source file: font.c
}

func (ec *execContext) internalCharFont_autogen(position, ch lispObject) (lispObject, error) {
	return ec.stub("internal-char-font") // Source file: font.c
}

func (ec *execContext) listFonts_autogen(font_spec, frame, num, prefer lispObject) (lispObject, error) {
	return ec.stub("list-fonts") // Source file: font.c
}

func (ec *execContext) openFont_autogen(font_entity, size, frame lispObject) (lispObject, error) {
	return ec.stub("open-font") // Source file: font.c
}

func (ec *execContext) queryFont_autogen(font_object lispObject) (lispObject, error) {
	return ec.stub("query-font") // Source file: font.c
}

func (ec *execContext) fontsetFont_autogen(name, ch, all lispObject) (lispObject, error) {
	return ec.stub("fontset-font") // Source file: fontset.c
}

func (ec *execContext) fontsetInfo_autogen(fontset, frame lispObject) (lispObject, error) {
	return ec.stub("fontset-info") // Source file: fontset.c
}

func (ec *execContext) fontsetList_autogen() (lispObject, error) {
	return ec.stub("fontset-list") // Source file: fontset.c
}

func (ec *execContext) fontsetListAll_autogen() (lispObject, error) {
	return ec.stub("fontset-list-all") // Source file: fontset.c
}

func (ec *execContext) newFontset_autogen(name, fontlist lispObject) (lispObject, error) {
	return ec.stub("new-fontset") // Source file: fontset.c
}

func (ec *execContext) queryFontset_autogen(pattern, regexpp lispObject) (lispObject, error) {
	return ec.stub("query-fontset") // Source file: fontset.c
}

func (ec *execContext) setFontsetFont_autogen(fontset, characters, font_spec, frame, add lispObject) (lispObject, error) {
	return ec.stub("set-fontset-font") // Source file: fontset.c
}

func (ec *execContext) deleteFrame_autogen(frame, force lispObject) (lispObject, error) {
	return ec.stub("delete-frame") // Source file: frame.c
}

func (ec *execContext) frameSetWasInvisible_autogen(frame, was_invisible lispObject) (lispObject, error) {
	return ec.stub("frame--set-was-invisible") // Source file: frame.c
}

func (ec *execContext) frameAfterMakeFrame_autogen(frame, made lispObject) (lispObject, error) {
	return ec.stub("frame-after-make-frame") // Source file: frame.c
}

func (ec *execContext) frameAncestorP_autogen(ancestor, descendant lispObject) (lispObject, error) {
	return ec.stub("frame-ancestor-p") // Source file: frame.c
}

func (ec *execContext) bottomDividerWidth_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("frame-bottom-divider-width") // Source file: frame.c
}

func (ec *execContext) frameCharHeight_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("frame-char-height") // Source file: frame.c
}

func (ec *execContext) frameCharWidth_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("frame-char-width") // Source file: frame.c
}

func (ec *execContext) frameChildFrameBorderWidth_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("frame-child-frame-border-width") // Source file: frame.c
}

func (ec *execContext) frameFocus_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("frame-focus") // Source file: frame.c
}

func (ec *execContext) fringeWidth_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("frame-fringe-width") // Source file: frame.c
}

func (ec *execContext) frameInternalBorderWidth_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("frame-internal-border-width") // Source file: frame.c
}

func (ec *execContext) frameList_autogen() (lispObject, error) {
	return ec.stub("frame-list") // Source file: frame.c
}

func (ec *execContext) frameLiveP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("frame-live-p") // Source file: frame.c
}

func (ec *execContext) frameNativeHeight_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("frame-native-height") // Source file: frame.c
}

func (ec *execContext) frameNativeWidth_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("frame-native-width") // Source file: frame.c
}

func (ec *execContext) frameParameter_autogen(frame, parameter lispObject) (lispObject, error) {
	return ec.stub("frame-parameter") // Source file: frame.c
}

func (ec *execContext) frameParameters_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("frame-parameters") // Source file: frame.c
}

func (ec *execContext) frameParent_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("frame-parent") // Source file: frame.c
}

func (ec *execContext) framePointerVisibleP_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("frame-pointer-visible-p") // Source file: frame.c
}

func (ec *execContext) framePosition_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("frame-position") // Source file: frame.c
}

func (ec *execContext) rightDividerWidth_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("frame-right-divider-width") // Source file: frame.c
}

func (ec *execContext) frameScaleFactor_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("frame-scale-factor") // Source file: frame.c
}

func (ec *execContext) scrollBarHeight_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("frame-scroll-bar-height") // Source file: frame.c
}

func (ec *execContext) scrollBarWidth_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("frame-scroll-bar-width") // Source file: frame.c
}

func (ec *execContext) frameTextCols_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("frame-text-cols") // Source file: frame.c
}

func (ec *execContext) frameTextHeight_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("frame-text-height") // Source file: frame.c
}

func (ec *execContext) frameTextLines_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("frame-text-lines") // Source file: frame.c
}

func (ec *execContext) frameTextWidth_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("frame-text-width") // Source file: frame.c
}

func (ec *execContext) frameTotalCols_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("frame-total-cols") // Source file: frame.c
}

func (ec *execContext) frameTotalLines_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("frame-total-lines") // Source file: frame.c
}

func (ec *execContext) frameVisibleP_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("frame-visible-p") // Source file: frame.c
}

func (ec *execContext) frameWindowStateChange_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("frame-window-state-change") // Source file: frame.c
}

func (ec *execContext) frameWindowsMinSize_autogen(frame, horizontal, ignore, pixelwise lispObject) (lispObject, error) {
	return ec.stub("frame-windows-min-size") // Source file: frame.c, attributes: const
}

func (ec *execContext) framep_autogen(object lispObject) (lispObject, error) {
	return ec.stub("framep") // Source file: frame.c
}

func (ec *execContext) handleSwitchFrame_autogen(event lispObject) (lispObject, error) {
	return ec.stub("handle-switch-frame") // Source file: frame.c
}

func (ec *execContext) iconifyFrame_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("iconify-frame") // Source file: frame.c
}

func (ec *execContext) lastNonminibufFrame_autogen() (lispObject, error) {
	return ec.stub("last-nonminibuffer-frame") // Source file: frame.c
}

func (ec *execContext) lowerFrame_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("lower-frame") // Source file: frame.c
}

func (ec *execContext) makeFrameInvisible_autogen(frame, force lispObject) (lispObject, error) {
	return ec.stub("make-frame-invisible") // Source file: frame.c
}

func (ec *execContext) makeFrameVisible_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("make-frame-visible") // Source file: frame.c
}

func (ec *execContext) makeTerminalFrame_autogen(parms lispObject) (lispObject, error) {
	return ec.stub("make-terminal-frame") // Source file: frame.c
}

func (ec *execContext) modifyFrameParameters_autogen(frame, alist lispObject) (lispObject, error) {
	return ec.stub("modify-frame-parameters") // Source file: frame.c
}

func (ec *execContext) mousePixelPosition_autogen() (lispObject, error) {
	return ec.stub("mouse-pixel-position") // Source file: frame.c
}

func (ec *execContext) mousePosition_autogen() (lispObject, error) {
	return ec.stub("mouse-position") // Source file: frame.c
}

func (ec *execContext) nextFrame_autogen(frame, miniframe lispObject) (lispObject, error) {
	return ec.stub("next-frame") // Source file: frame.c
}

func (ec *execContext) oldSelectedFrame_autogen() (lispObject, error) {
	return ec.stub("old-selected-frame") // Source file: frame.c
}

func (ec *execContext) previousFrame_autogen(frame, miniframe lispObject) (lispObject, error) {
	return ec.stub("previous-frame") // Source file: frame.c
}

func (ec *execContext) raiseFrame_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("raise-frame") // Source file: frame.c
}

func (ec *execContext) reconsiderFrameFonts_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("reconsider-frame-fonts") // Source file: frame.c
}

func (ec *execContext) redirectFrameFocus_autogen(frame, focus_frame lispObject) (lispObject, error) {
	return ec.stub("redirect-frame-focus") // Source file: frame.c
}

func (ec *execContext) selectFrame_autogen(frame, norecord lispObject) (lispObject, error) {
	return ec.stub("select-frame") // Source file: frame.c
}

func (ec *execContext) selectedFrame_autogen() (lispObject, error) {
	return ec.stub("selected-frame") // Source file: frame.c
}

func (ec *execContext) setFrameHeight_autogen(frame, height, pretend, pixelwise lispObject) (lispObject, error) {
	return ec.stub("set-frame-height") // Source file: frame.c
}

func (ec *execContext) setFramePosition_autogen(frame, x, y lispObject) (lispObject, error) {
	return ec.stub("set-frame-position") // Source file: frame.c
}

func (ec *execContext) setFrameSize_autogen(frame, width, height, pixelwise lispObject) (lispObject, error) {
	return ec.stub("set-frame-size") // Source file: frame.c
}

func (ec *execContext) setFrameWidth_autogen(frame, width, pretend, pixelwise lispObject) (lispObject, error) {
	return ec.stub("set-frame-width") // Source file: frame.c
}

func (ec *execContext) setFrameWindowStateChange_autogen(frame, arg lispObject) (lispObject, error) {
	return ec.stub("set-frame-window-state-change") // Source file: frame.c
}

func (ec *execContext) setMousePixelPosition_autogen(frame, x, y lispObject) (lispObject, error) {
	return ec.stub("set-mouse-pixel-position") // Source file: frame.c
}

func (ec *execContext) setMousePosition_autogen(frame, x, y lispObject) (lispObject, error) {
	return ec.stub("set-mouse-position") // Source file: frame.c
}

func (ec *execContext) toolBarPixelWidth_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("tool-bar-pixel-width") // Source file: frame.c
}

func (ec *execContext) visibleFrameList_autogen() (lispObject, error) {
	return ec.stub("visible-frame-list") // Source file: frame.c
}

func (ec *execContext) windowSystem_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("window-system") // Source file: frame.c
}

func (ec *execContext) xFocusFrame_autogen(frame, noactivate lispObject) (lispObject, error) {
	return ec.stub("x-focus-frame") // Source file: frame.c
}

func (ec *execContext) xGetResource_autogen(attribute, class, component, subclass lispObject) (lispObject, error) {
	return ec.stub("x-get-resource") // Source file: frame.c
}

func (ec *execContext) xParseGeometry_autogen(string lispObject) (lispObject, error) {
	return ec.stub("x-parse-geometry") // Source file: frame.c
}

func (ec *execContext) defineFringeBitmap_autogen(bitmap, bits, height, width, align lispObject) (lispObject, error) {
	return ec.stub("define-fringe-bitmap") // Source file: fringe.c
}

func (ec *execContext) destroyFringeBitmap_autogen(bitmap lispObject) (lispObject, error) {
	return ec.stub("destroy-fringe-bitmap") // Source file: fringe.c
}

func (ec *execContext) fringeBitmapsAtPos_autogen(pos, window lispObject) (lispObject, error) {
	return ec.stub("fringe-bitmaps-at-pos") // Source file: fringe.c
}

func (ec *execContext) setFringeBitmapFace_autogen(bitmap, face lispObject) (lispObject, error) {
	return ec.stub("set-fringe-bitmap-face") // Source file: fringe.c
}

func (ec *execContext) gfileAddWatch_autogen(file, flags, callback lispObject) (lispObject, error) {
	return ec.stub("gfile-add-watch") // Source file: gfilenotify.c
}

func (ec *execContext) gfileMonitorName_autogen(watch_descriptor lispObject) (lispObject, error) {
	return ec.stub("gfile-monitor-name") // Source file: gfilenotify.c
}

func (ec *execContext) gfileRmWatch_autogen(watch_descriptor lispObject) (lispObject, error) {
	return ec.stub("gfile-rm-watch") // Source file: gfilenotify.c
}

func (ec *execContext) gfileValidP_autogen(watch_descriptor lispObject) (lispObject, error) {
	return ec.stub("gfile-valid-p") // Source file: gfilenotify.c
}

func (ec *execContext) gnutlsAsynchronousParameters_autogen(proc, params lispObject) (lispObject, error) {
	return ec.stub("gnutls-asynchronous-parameters") // Source file: gnutls.c
}

func (ec *execContext) gnutlsAvailableP_autogen() (lispObject, error) {
	return ec.stub("gnutls-available-p") // Source file: gnutls.c
}

func (ec *execContext) gnutlsBoot_autogen(proc, type_, proplist lispObject) (lispObject, error) {
	return ec.stub("gnutls-boot") // Source file: gnutls.c
}

func (ec *execContext) gnutlsBye_autogen(proc, cont lispObject) (lispObject, error) {
	return ec.stub("gnutls-bye") // Source file: gnutls.c
}

func (ec *execContext) gnutlsCiphers_autogen() (lispObject, error) {
	return ec.stub("gnutls-ciphers") // Source file: gnutls.c
}

func (ec *execContext) gnutlsDeinit_autogen(proc lispObject) (lispObject, error) {
	return ec.stub("gnutls-deinit") // Source file: gnutls.c
}

func (ec *execContext) gnutlsDigests_autogen() (lispObject, error) {
	return ec.stub("gnutls-digests") // Source file: gnutls.c
}

func (ec *execContext) gnutlsErrorFatalp_autogen(err lispObject) (lispObject, error) {
	return ec.stub("gnutls-error-fatalp") // Source file: gnutls.c
}

func (ec *execContext) gnutlsErrorString_autogen(err lispObject) (lispObject, error) {
	return ec.stub("gnutls-error-string") // Source file: gnutls.c
}

func (ec *execContext) gnutlsErrorp_autogen(err lispObject) (lispObject, error) {
	return ec.stub("gnutls-errorp") // Source file: gnutls.c, attributes: const
}

func (ec *execContext) gnutlsFormatCertificate_autogen(cert lispObject) (lispObject, error) {
	return ec.stub("gnutls-format-certificate") // Source file: gnutls.c
}

func (ec *execContext) gnutlsGetInitstage_autogen(proc lispObject) (lispObject, error) {
	return ec.stub("gnutls-get-initstage") // Source file: gnutls.c
}

func (ec *execContext) gnutlsHashDigest_autogen(digest_method, input lispObject) (lispObject, error) {
	return ec.stub("gnutls-hash-digest") // Source file: gnutls.c
}

func (ec *execContext) gnutlsHashMac_autogen(hash_method, key, input lispObject) (lispObject, error) {
	return ec.stub("gnutls-hash-mac") // Source file: gnutls.c
}

func (ec *execContext) gnutlsMacs_autogen() (lispObject, error) {
	return ec.stub("gnutls-macs") // Source file: gnutls.c
}

func (ec *execContext) gnutlsPeerStatus_autogen(proc lispObject) (lispObject, error) {
	return ec.stub("gnutls-peer-status") // Source file: gnutls.c
}

func (ec *execContext) gnutlsPeerStatusWarningDescribe_autogen(status_symbol lispObject) (lispObject, error) {
	return ec.stub("gnutls-peer-status-warning-describe") // Source file: gnutls.c
}

func (ec *execContext) gnutlsSymmetricDecrypt_autogen(cipher, key, iv, input, aead_auth lispObject) (lispObject, error) {
	return ec.stub("gnutls-symmetric-decrypt") // Source file: gnutls.c
}

func (ec *execContext) gnutlsSymmetricEncrypt_autogen(cipher, key, iv, input, aead_auth lispObject) (lispObject, error) {
	return ec.stub("gnutls-symmetric-encrypt") // Source file: gnutls.c
}

func (ec *execContext) haikuDisplayMonitorAttributesList_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("haiku-display-monitor-attributes-list") // Source file: haikufns.c
}

func (ec *execContext) haikuFrameEdges_autogen(frame, type_ lispObject) (lispObject, error) {
	return ec.stub("haiku-frame-edges") // Source file: haikufns.c
}

func (ec *execContext) haikuFrameGeometry_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("haiku-frame-geometry") // Source file: haikufns.c
}

func (ec *execContext) haikuFrameListZOrder_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("haiku-frame-list-z-order") // Source file: haikufns.c
}

func (ec *execContext) haikuFrameRestack_autogen(frame1, frame2, above lispObject) (lispObject, error) {
	return ec.stub("haiku-frame-restack") // Source file: haikufns.c
}

func (ec *execContext) haikuGetVersionString_autogen() (lispObject, error) {
	return ec.stub("haiku-get-version-string") // Source file: haikufns.c
}

func (ec *execContext) haikuMouseAbsolutePixelPosition_autogen() (lispObject, error) {
	return ec.stub("haiku-mouse-absolute-pixel-position") // Source file: haikufns.c
}

func (ec *execContext) haikuPutResource_autogen(resource, string lispObject) (lispObject, error) {
	return ec.stub("haiku-put-resource") // Source file: haikufns.c
}

func (ec *execContext) haikuReadFileName_autogen(prompt, frame, dir, mustmatch, dir_only_p, save_text lispObject) (lispObject, error) {
	return ec.stub("haiku-read-file-name") // Source file: haikufns.c
}

func (ec *execContext) haikuSaveSessionReply_autogen(quit_reply lispObject) (lispObject, error) {
	return ec.stub("haiku-save-session-reply") // Source file: haikufns.c
}

func (ec *execContext) haikuSetMouseAbsolutePixelPosition_autogen(x, y lispObject) (lispObject, error) {
	return ec.stub("haiku-set-mouse-absolute-pixel-position") // Source file: haikufns.c
}

func (ec *execContext) xCloseConnection_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-close-connection") // Source file: haikufns.c, attributes: noreturn
}

func (ec *execContext) xCloseConnection_1_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-close-connection") // Source file: pgtkfns.c
}

func (ec *execContext) xCloseConnection_2_autogen(display lispObject) (lispObject, error) {
	return ec.stub("x-close-connection") // Source file: w32fns.c
}

func (ec *execContext) xCloseConnection_3_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-close-connection") // Source file: xfns.c
}

func (ec *execContext) xCreateFrame_autogen(parms lispObject) (lispObject, error) {
	return ec.stub("x-create-frame") // Source file: haikufns.c
}

func (ec *execContext) xCreateFrame_1_autogen(parms lispObject) (lispObject, error) {
	return ec.stub("x-create-frame") // Source file: pgtkfns.c
}

func (ec *execContext) xCreateFrame_2_autogen(parameters lispObject) (lispObject, error) {
	return ec.stub("x-create-frame") // Source file: w32fns.c
}

func (ec *execContext) xCreateFrame_3_autogen(parms lispObject) (lispObject, error) {
	return ec.stub("x-create-frame") // Source file: xfns.c
}

func (ec *execContext) xDisplayBackingStore_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-backing-store") // Source file: haikufns.c
}

func (ec *execContext) xDisplayBackingStore_1_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-backing-store") // Source file: pgtkfns.c
}

func (ec *execContext) xDisplayBackingStore_2_autogen(display lispObject) (lispObject, error) {
	return ec.stub("x-display-backing-store") // Source file: w32fns.c
}

func (ec *execContext) xDisplayBackingStore_3_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-backing-store") // Source file: xfns.c
}

func (ec *execContext) xDisplayColorCells_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-color-cells") // Source file: haikufns.c
}

func (ec *execContext) xDisplayColorCells_1_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-color-cells") // Source file: pgtkfns.c
}

func (ec *execContext) xDisplayColorCells_2_autogen(display lispObject) (lispObject, error) {
	return ec.stub("x-display-color-cells") // Source file: w32fns.c
}

func (ec *execContext) xDisplayColorCells_3_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-color-cells") // Source file: xfns.c
}

func (ec *execContext) xDisplayGrayscaleP_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-grayscale-p") // Source file: haikufns.c
}

func (ec *execContext) xDisplayGrayscaleP_1_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-grayscale-p") // Source file: pgtkfns.c
}

func (ec *execContext) xDisplayGrayscaleP_2_autogen(display lispObject) (lispObject, error) {
	return ec.stub("x-display-grayscale-p") // Source file: w32fns.c
}

func (ec *execContext) xDisplayGrayscaleP_3_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-grayscale-p") // Source file: xfns.c
}

func (ec *execContext) xDisplayList_autogen() (lispObject, error) {
	return ec.stub("x-display-list") // Source file: haikufns.c
}

func (ec *execContext) xDisplayList_1_autogen() (lispObject, error) {
	return ec.stub("x-display-list") // Source file: pgtkfns.c
}

func (ec *execContext) xDisplayList_2_autogen() (lispObject, error) {
	return ec.stub("x-display-list") // Source file: w32fns.c
}

func (ec *execContext) xDisplayList_3_autogen() (lispObject, error) {
	return ec.stub("x-display-list") // Source file: xfns.c
}

func (ec *execContext) xDisplayMmHeight_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-mm-height") // Source file: haikufns.c
}

func (ec *execContext) xDisplayMmHeight_1_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-mm-height") // Source file: pgtkfns.c
}

func (ec *execContext) xDisplayMmHeight_2_autogen(display lispObject) (lispObject, error) {
	return ec.stub("x-display-mm-height") // Source file: w32fns.c
}

func (ec *execContext) xDisplayMmHeight_3_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-mm-height") // Source file: xfns.c
}

func (ec *execContext) xDisplayMmWidth_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-mm-width") // Source file: haikufns.c
}

func (ec *execContext) xDisplayMmWidth_1_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-mm-width") // Source file: pgtkfns.c
}

func (ec *execContext) xDisplayMmWidth_2_autogen(display lispObject) (lispObject, error) {
	return ec.stub("x-display-mm-width") // Source file: w32fns.c
}

func (ec *execContext) xDisplayMmWidth_3_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-mm-width") // Source file: xfns.c
}

func (ec *execContext) xDisplayPixelHeight_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-pixel-height") // Source file: haikufns.c
}

func (ec *execContext) xDisplayPixelHeight_1_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-pixel-height") // Source file: pgtkfns.c
}

func (ec *execContext) xDisplayPixelHeight_2_autogen(display lispObject) (lispObject, error) {
	return ec.stub("x-display-pixel-height") // Source file: w32fns.c
}

func (ec *execContext) xDisplayPixelHeight_3_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-pixel-height") // Source file: xfns.c
}

func (ec *execContext) xDisplayPixelWidth_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-pixel-width") // Source file: haikufns.c
}

func (ec *execContext) xDisplayPixelWidth_1_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-pixel-width") // Source file: pgtkfns.c
}

func (ec *execContext) xDisplayPixelWidth_2_autogen(display lispObject) (lispObject, error) {
	return ec.stub("x-display-pixel-width") // Source file: w32fns.c
}

func (ec *execContext) xDisplayPixelWidth_3_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-pixel-width") // Source file: xfns.c
}

func (ec *execContext) xDisplayPlanes_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-planes") // Source file: haikufns.c
}

func (ec *execContext) xDisplayPlanes_1_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-planes") // Source file: pgtkfns.c
}

func (ec *execContext) xDisplayPlanes_2_autogen(display lispObject) (lispObject, error) {
	return ec.stub("x-display-planes") // Source file: w32fns.c
}

func (ec *execContext) xDisplayPlanes_3_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-planes") // Source file: xfns.c
}

func (ec *execContext) xDisplaySaveUnder_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-save-under") // Source file: haikufns.c
}

func (ec *execContext) xDisplaySaveUnder_1_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-save-under") // Source file: pgtkfns.c
}

func (ec *execContext) xDisplaySaveUnder_2_autogen(display lispObject) (lispObject, error) {
	return ec.stub("x-display-save-under") // Source file: w32fns.c
}

func (ec *execContext) xDisplaySaveUnder_3_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-save-under") // Source file: xfns.c
}

func (ec *execContext) xDisplayScreens_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-screens") // Source file: haikufns.c
}

func (ec *execContext) xDisplayScreens_1_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-screens") // Source file: pgtkfns.c
}

func (ec *execContext) xDisplayScreens_2_autogen(display lispObject) (lispObject, error) {
	return ec.stub("x-display-screens") // Source file: w32fns.c
}

func (ec *execContext) xDisplayScreens_3_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-screens") // Source file: xfns.c
}

func (ec *execContext) xDisplayVisualClass_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-visual-class") // Source file: haikufns.c
}

func (ec *execContext) xDisplayVisualClass_1_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-visual-class") // Source file: pgtkfns.c
}

func (ec *execContext) xDisplayVisualClass_2_autogen(display lispObject) (lispObject, error) {
	return ec.stub("x-display-visual-class") // Source file: w32fns.c
}

func (ec *execContext) xDisplayVisualClass_3_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-visual-class") // Source file: xfns.c
}

func (ec *execContext) xDoubleBufferedP_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("x-double-buffered-p") // Source file: haikufns.c
}

func (ec *execContext) xDoubleBufferedP_1_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("x-double-buffered-p") // Source file: xfns.c
}

func (ec *execContext) xHideTip_autogen() (lispObject, error) {
	return ec.stub("x-hide-tip") // Source file: haikufns.c
}

func (ec *execContext) xHideTip_1_autogen() (lispObject, error) {
	return ec.stub("x-hide-tip") // Source file: pgtkfns.c
}

func (ec *execContext) xHideTip_2_autogen() (lispObject, error) {
	return ec.stub("x-hide-tip") // Source file: w32fns.c
}

func (ec *execContext) xHideTip_3_autogen() (lispObject, error) {
	return ec.stub("x-hide-tip") // Source file: xfns.c
}

func (ec *execContext) xOpenConnection_autogen(display, resource_string, must_succeed lispObject) (lispObject, error) {
	return ec.stub("x-open-connection") // Source file: haikufns.c
}

func (ec *execContext) xOpenConnection_1_autogen(display, resource_string, must_succeed lispObject) (lispObject, error) {
	return ec.stub("x-open-connection") // Source file: pgtkfns.c
}

func (ec *execContext) xOpenConnection_2_autogen(display, xrm_string, must_succeed lispObject) (lispObject, error) {
	return ec.stub("x-open-connection") // Source file: w32fns.c
}

func (ec *execContext) xOpenConnection_3_autogen(display, xrm_string, must_succeed lispObject) (lispObject, error) {
	return ec.stub("x-open-connection") // Source file: xfns.c
}

func (ec *execContext) xServerVendor_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-server-vendor") // Source file: haikufns.c
}

func (ec *execContext) xServerVendor_1_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-server-vendor") // Source file: w32fns.c
}

func (ec *execContext) xServerVendor_2_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-server-vendor") // Source file: xfns.c
}

func (ec *execContext) xServerVersion_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-server-version") // Source file: haikufns.c
}

func (ec *execContext) xServerVersion_1_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-server-version") // Source file: w32fns.c
}

func (ec *execContext) xServerVersion_2_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-server-version") // Source file: xfns.c
}

func (ec *execContext) xShowTip_autogen(string, frame, parms, timeout, dx, dy lispObject) (lispObject, error) {
	return ec.stub("x-show-tip") // Source file: haikufns.c
}

func (ec *execContext) xShowTip_1_autogen(string, frame, parms, timeout, dx, dy lispObject) (lispObject, error) {
	return ec.stub("x-show-tip") // Source file: pgtkfns.c
}

func (ec *execContext) xShowTip_2_autogen(string, frame, parms, timeout, dx, dy lispObject) (lispObject, error) {
	return ec.stub("x-show-tip") // Source file: w32fns.c
}

func (ec *execContext) xShowTip_3_autogen(string, frame, parms, timeout, dx, dy lispObject) (lispObject, error) {
	return ec.stub("x-show-tip") // Source file: xfns.c
}

func (ec *execContext) xwColorDefinedP_autogen(color, frame lispObject) (lispObject, error) {
	return ec.stub("xw-color-defined-p") // Source file: haikufns.c
}

func (ec *execContext) xwColorDefinedP_1_autogen(color, frame lispObject) (lispObject, error) {
	return ec.stub("xw-color-defined-p") // Source file: pgtkfns.c
}

func (ec *execContext) xwColorDefinedP_2_autogen(color, frame lispObject) (lispObject, error) {
	return ec.stub("xw-color-defined-p") // Source file: w32fns.c
}

func (ec *execContext) xwColorDefinedP_3_autogen(color, frame lispObject) (lispObject, error) {
	return ec.stub("xw-color-defined-p") // Source file: xfns.c
}

func (ec *execContext) xwColorValues_autogen(color, frame lispObject) (lispObject, error) {
	return ec.stub("xw-color-values") // Source file: haikufns.c
}

func (ec *execContext) xwColorValues_1_autogen(color, frame lispObject) (lispObject, error) {
	return ec.stub("xw-color-values") // Source file: pgtkfns.c
}

func (ec *execContext) xwColorValues_2_autogen(color, frame lispObject) (lispObject, error) {
	return ec.stub("xw-color-values") // Source file: w32fns.c
}

func (ec *execContext) xwColorValues_3_autogen(color, frame lispObject) (lispObject, error) {
	return ec.stub("xw-color-values") // Source file: xfns.c
}

func (ec *execContext) xwDisplayColorP_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("xw-display-color-p") // Source file: haikufns.c
}

func (ec *execContext) xwDisplayColorP_1_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("xw-display-color-p") // Source file: pgtkfns.c
}

func (ec *execContext) xwDisplayColorP_2_autogen(display lispObject) (lispObject, error) {
	return ec.stub("xw-display-color-p") // Source file: w32fns.c
}

func (ec *execContext) xwDisplayColorP_3_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("xw-display-color-p") // Source file: xfns.c
}

func (ec *execContext) fontGetSystemFont_autogen() (lispObject, error) {
	return ec.stub("font-get-system-font") // Source file: haikufont.c
}

func (ec *execContext) fontGetSystemFont_1_autogen() (lispObject, error) {
	return ec.stub("font-get-system-font") // Source file: xsettings.c
}

func (ec *execContext) fontGetSystemNormalFont_autogen() (lispObject, error) {
	return ec.stub("font-get-system-normal-font") // Source file: haikufont.c
}

func (ec *execContext) fontGetSystemNormalFont_1_autogen() (lispObject, error) {
	return ec.stub("font-get-system-normal-font") // Source file: xsettings.c
}

func (ec *execContext) xSelectFont_autogen(frame, exclude_proportional lispObject) (lispObject, error) {
	return ec.stub("x-select-font") // Source file: haikufont.c
}

func (ec *execContext) xSelectFont_1_autogen(frame, ignored lispObject) (lispObject, error) {
	return ec.stub("x-select-font") // Source file: pgtkfns.c
}

func (ec *execContext) xSelectFont_2_autogen(frame, exclude_proportional lispObject) (lispObject, error) {
	return ec.stub("x-select-font") // Source file: w32font.c
}

func (ec *execContext) xSelectFont_3_autogen(frame, ignored lispObject) (lispObject, error) {
	return ec.stub("x-select-font") // Source file: xfns.c
}

func (ec *execContext) haikuMenuBarOpen_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("haiku-menu-bar-open") // Source file: haikumenu.c
}

func (ec *execContext) menuOrPopupActiveP_autogen() (lispObject, error) {
	return ec.stub("menu-or-popup-active-p") // Source file: haikumenu.c
}

func (ec *execContext) menuOrPopupActiveP_1_autogen() (lispObject, error) {
	return ec.stub("menu-or-popup-active-p") // Source file: pgtkmenu.c
}

func (ec *execContext) menuOrPopupActiveP_2_autogen() (lispObject, error) {
	return ec.stub("menu-or-popup-active-p") // Source file: w32menu.c
}

func (ec *execContext) menuOrPopupActiveP_3_autogen() (lispObject, error) {
	return ec.stub("menu-or-popup-active-p") // Source file: xmenu.c
}

func (ec *execContext) haikuDragMessage_autogen(frame, message, allow_same_frame, follow_tooltip lispObject) (lispObject, error) {
	return ec.stub("haiku-drag-message") // Source file: haikuselect.c
}

func (ec *execContext) haikuRosterLaunch_autogen(file_or_type, args lispObject) (lispObject, error) {
	return ec.stub("haiku-roster-launch") // Source file: haikuselect.c
}

func (ec *execContext) haikuSelectionData_autogen(clipboard, name lispObject) (lispObject, error) {
	return ec.stub("haiku-selection-data") // Source file: haikuselect.c
}

func (ec *execContext) haikuSelectionOwnerP_autogen(selection lispObject) (lispObject, error) {
	return ec.stub("haiku-selection-owner-p") // Source file: haikuselect.c
}

func (ec *execContext) haikuSelectionPut_autogen(clipboard, name, data, clear lispObject) (lispObject, error) {
	return ec.stub("haiku-selection-put") // Source file: haikuselect.c
}

func (ec *execContext) haikuSelectionTimestamp_autogen(clipboard lispObject) (lispObject, error) {
	return ec.stub("haiku-selection-timestamp") // Source file: haikuselect.c
}

func (ec *execContext) haikuSendMessage_autogen(program, message lispObject) (lispObject, error) {
	return ec.stub("haiku-send-message") // Source file: haikuselect.c
}

func (ec *execContext) haikuWriteNodeAttribute_autogen(file, name, message lispObject) (lispObject, error) {
	return ec.stub("haiku-write-node-attribute") // Source file: haikuselect.c
}

func (ec *execContext) clearImageCache_autogen(filter, animation_cache lispObject) (lispObject, error) {
	return ec.stub("clear-image-cache") // Source file: image.c
}

func (ec *execContext) imageCacheSize_autogen() (lispObject, error) {
	return ec.stub("image-cache-size") // Source file: image.c
}

func (ec *execContext) imageFlush_autogen(spec, frame lispObject) (lispObject, error) {
	return ec.stub("image-flush") // Source file: image.c
}

func (ec *execContext) imageMaskP_autogen(spec, frame lispObject) (lispObject, error) {
	return ec.stub("image-mask-p") // Source file: image.c
}

func (ec *execContext) imageMetadata_autogen(spec, frame lispObject) (lispObject, error) {
	return ec.stub("image-metadata") // Source file: image.c
}

func (ec *execContext) imageSize_autogen(spec, pixels, frame lispObject) (lispObject, error) {
	return ec.stub("image-size") // Source file: image.c
}

func (ec *execContext) imageTransformsP_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("image-transforms-p") // Source file: image.c
}

func (ec *execContext) imagemagickTypes_autogen() (lispObject, error) {
	return ec.stub("imagemagick-types") // Source file: image.c
}

func (ec *execContext) imagep_autogen(spec lispObject) (lispObject, error) {
	return ec.stub("imagep") // Source file: image.c
}

func (ec *execContext) initImageLibrary_autogen(type_ lispObject) (lispObject, error) {
	return ec.stub("init-image-library") // Source file: image.c
}

func (ec *execContext) lookupImage_autogen(spec lispObject) (lispObject, error) {
	return ec.stub("lookup-image") // Source file: image.c
}

func (ec *execContext) computeMotion_autogen(from, frompos, to, topos, width, offsets, window lispObject) (lispObject, error) {
	return ec.stub("compute-motion") // Source file: indent.c
}

func (ec *execContext) currentColumn_autogen() (lispObject, error) {
	return ec.stub("current-column") // Source file: indent.c
}

func (ec *execContext) currentIndentation_autogen() (lispObject, error) {
	return ec.stub("current-indentation") // Source file: indent.c
}

func (ec *execContext) indentTo_autogen(column, minimum lispObject) (lispObject, error) {
	return ec.stub("indent-to") // Source file: indent.c
}

func (ec *execContext) lineNumberDisplayWidth_autogen(pixelwise lispObject) (lispObject, error) {
	return ec.stub("line-number-display-width") // Source file: indent.c
}

func (ec *execContext) moveToColumn_autogen(column, force lispObject) (lispObject, error) {
	return ec.stub("move-to-column") // Source file: indent.c
}

func (ec *execContext) verticalMotion_autogen(lines, window, cur_col lispObject) (lispObject, error) {
	return ec.stub("vertical-motion") // Source file: indent.c
}

func (ec *execContext) inotifyAddWatch_autogen(filename, aspect, callback lispObject) (lispObject, error) {
	return ec.stub("inotify-add-watch") // Source file: inotify.c
}

func (ec *execContext) inotifyAllocatedP_autogen() (lispObject, error) {
	return ec.stub("inotify-allocated-p") // Source file: inotify.c
}

func (ec *execContext) inotifyRmWatch_autogen(watch_descriptor lispObject) (lispObject, error) {
	return ec.stub("inotify-rm-watch") // Source file: inotify.c
}

func (ec *execContext) inotifyValidP_autogen(watch_descriptor lispObject) (lispObject, error) {
	return ec.stub("inotify-valid-p") // Source file: inotify.c
}

func (ec *execContext) inotifyWatchList_autogen() (lispObject, error) {
	return ec.stub("inotify-watch-list") // Source file: inotify.c
}

func (ec *execContext) combineAfterChangeExecute_autogen() (lispObject, error) {
	return ec.stub("combine-after-change-execute") // Source file: insdel.c
}

func (ec *execContext) jsonAvailableP_autogen() (lispObject, error) {
	return ec.stub("json--available-p") // Source file: json.c
}

func (ec *execContext) jsonInsert_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("json-insert") // Source file: json.c
}

func (ec *execContext) jsonParseBuffer_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("json-parse-buffer") // Source file: json.c
}

func (ec *execContext) jsonParseString_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("json-parse-string") // Source file: json.c
}

func (ec *execContext) jsonSerialize_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("json-serialize") // Source file: json.c
}

func (ec *execContext) abortRecursiveEdit_autogen() (lispObject, error) {
	return ec.stub("abort-recursive-edit") // Source file: keyboard.c, attributes: noreturn
}

func (ec *execContext) clearThisCommandKeys_autogen(keep_record lispObject) (lispObject, error) {
	return ec.stub("clear-this-command-keys") // Source file: keyboard.c
}

func (ec *execContext) commandErrorDefaultFunction_autogen(data, context, signal lispObject) (lispObject, error) {
	return ec.stub("command-error-default-function") // Source file: keyboard.c
}

func (ec *execContext) currentIdleTime_autogen() (lispObject, error) {
	return ec.stub("current-idle-time") // Source file: keyboard.c
}

func (ec *execContext) currentInputMode_autogen() (lispObject, error) {
	return ec.stub("current-input-mode") // Source file: keyboard.c
}

func (ec *execContext) discardInput_autogen() (lispObject, error) {
	return ec.stub("discard-input") // Source file: keyboard.c
}

func (ec *execContext) eventConvertList_autogen(event_desc lispObject) (lispObject, error) {
	return ec.stub("event-convert-list") // Source file: keyboard.c
}

func (ec *execContext) exitRecursiveEdit_autogen() (lispObject, error) {
	return ec.stub("exit-recursive-edit") // Source file: keyboard.c, attributes: noreturn
}

func (ec *execContext) inputPendingP_autogen(check_timers lispObject) (lispObject, error) {
	return ec.stub("input-pending-p") // Source file: keyboard.c
}

func (ec *execContext) internalTrackMouse_autogen(bodyfun lispObject) (lispObject, error) {
	return ec.stub("internal--track-mouse") // Source file: keyboard.c
}

func (ec *execContext) eventSymbolParseModifiers_autogen(symbol lispObject) (lispObject, error) {
	return ec.stub("internal-event-symbol-parse-modifiers") // Source file: keyboard.c
}

func (ec *execContext) internalHandleFocusIn_autogen(event lispObject) (lispObject, error) {
	return ec.stub("internal-handle-focus-in") // Source file: keyboard.c
}

func (ec *execContext) lossageSize_autogen(arg lispObject) (lispObject, error) {
	return ec.stub("lossage-size") // Source file: keyboard.c
}

func (ec *execContext) openDribbleFile_autogen(file lispObject) (lispObject, error) {
	return ec.stub("open-dribble-file") // Source file: keyboard.c
}

func (ec *execContext) posnAtPoint_autogen(pos, window lispObject) (lispObject, error) {
	return ec.stub("posn-at-point") // Source file: keyboard.c
}

func (ec *execContext) posnAtXY_autogen(x, y, frame_or_window, whole lispObject) (lispObject, error) {
	return ec.stub("posn-at-x-y") // Source file: keyboard.c
}

func (ec *execContext) readKeySequence_autogen(prompt, continue_echo, dont_downcase_last, can_return_switch_frame, cmd_loop lispObject) (lispObject, error) {
	return ec.stub("read-key-sequence") // Source file: keyboard.c
}

func (ec *execContext) readKeySequenceVector_autogen(prompt, continue_echo, dont_downcase_last, can_return_switch_frame, cmd_loop lispObject) (lispObject, error) {
	return ec.stub("read-key-sequence-vector") // Source file: keyboard.c
}

func (ec *execContext) recentKeys_autogen(include_cmds lispObject) (lispObject, error) {
	return ec.stub("recent-keys") // Source file: keyboard.c
}

func (ec *execContext) recursionDepth_autogen() (lispObject, error) {
	return ec.stub("recursion-depth") // Source file: keyboard.c
}

func (ec *execContext) recursiveEdit_autogen() (lispObject, error) {
	return ec.stub("recursive-edit") // Source file: keyboard.c
}

func (ec *execContext) setThisCommandKeys_autogen(keys lispObject) (lispObject, error) {
	return ec.stub("set--this-command-keys") // Source file: keyboard.c
}

func (ec *execContext) setInputInterruptMode_autogen(interrupt lispObject) (lispObject, error) {
	return ec.stub("set-input-interrupt-mode") // Source file: keyboard.c
}

func (ec *execContext) setInputMetaMode_autogen(meta, terminal lispObject) (lispObject, error) {
	return ec.stub("set-input-meta-mode") // Source file: keyboard.c
}

func (ec *execContext) setInputMode_autogen(interrupt, flow, meta, quit lispObject) (lispObject, error) {
	return ec.stub("set-input-mode") // Source file: keyboard.c
}

func (ec *execContext) setOutputFlowControl_autogen(flow, terminal lispObject) (lispObject, error) {
	return ec.stub("set-output-flow-control") // Source file: keyboard.c
}

func (ec *execContext) setQuitChar_autogen(quit lispObject) (lispObject, error) {
	return ec.stub("set-quit-char") // Source file: keyboard.c
}

func (ec *execContext) suspendEmacs_autogen(stuffstring lispObject) (lispObject, error) {
	return ec.stub("suspend-emacs") // Source file: keyboard.c
}

func (ec *execContext) thisCommandKeys_autogen() (lispObject, error) {
	return ec.stub("this-command-keys") // Source file: keyboard.c
}

func (ec *execContext) thisCommandKeysVector_autogen() (lispObject, error) {
	return ec.stub("this-command-keys-vector") // Source file: keyboard.c
}

func (ec *execContext) thisSingleCommandKeys_autogen() (lispObject, error) {
	return ec.stub("this-single-command-keys") // Source file: keyboard.c
}

func (ec *execContext) thisSingleCommandRawKeys_autogen() (lispObject, error) {
	return ec.stub("this-single-command-raw-keys") // Source file: keyboard.c
}

func (ec *execContext) topLevel_autogen() (lispObject, error) {
	return ec.stub("top-level") // Source file: keyboard.c, attributes: noreturn
}

func (ec *execContext) accessibleKeymaps_autogen(keymap, prefix lispObject) (lispObject, error) {
	return ec.stub("accessible-keymaps") // Source file: keymap.c
}

func (ec *execContext) commandRemapping_autogen(command, position, keymaps lispObject) (lispObject, error) {
	return ec.stub("command-remapping") // Source file: keymap.c
}

func (ec *execContext) copyKeymap_autogen(keymap lispObject) (lispObject, error) {
	return ec.stub("copy-keymap") // Source file: keymap.c
}

func (ec *execContext) currentActiveMaps_autogen(olp, position lispObject) (lispObject, error) {
	return ec.stub("current-active-maps") // Source file: keymap.c
}

func (ec *execContext) currentGlobalMap_autogen() (lispObject, error) {
	return ec.stub("current-global-map") // Source file: keymap.c
}

func (ec *execContext) currentLocalMap_autogen() (lispObject, error) {
	return ec.stub("current-local-map") // Source file: keymap.c
}

func (ec *execContext) currentMinorModeMaps_autogen() (lispObject, error) {
	return ec.stub("current-minor-mode-maps") // Source file: keymap.c
}

func (ec *execContext) defineKey_autogen(keymap, key, def, remove lispObject) (lispObject, error) {
	return ec.stub("define-key") // Source file: keymap.c
}

func (ec *execContext) describeBufferBindings_autogen(buffer, prefix, menus lispObject) (lispObject, error) {
	return ec.stub("describe-buffer-bindings") // Source file: keymap.c
}

func (ec *execContext) describeVector_autogen(vector, describer lispObject) (lispObject, error) {
	return ec.stub("describe-vector") // Source file: keymap.c
}

func (ec *execContext) helpDescribeVector_autogen(vector, prefix, describer, partial, shadow, entire_map, mention_shadow lispObject) (lispObject, error) {
	return ec.stub("help--describe-vector") // Source file: keymap.c
}

func (ec *execContext) keyBinding_autogen(key, accept_default, no_remap, position lispObject) (lispObject, error) {
	return ec.stub("key-binding") // Source file: keymap.c
}

func (ec *execContext) keyDescription_autogen(keys, prefix lispObject) (lispObject, error) {
	return ec.stub("key-description") // Source file: keymap.c
}

func (ec *execContext) keymapGetKeyelt_autogen(object, autoload lispObject) (lispObject, error) {
	return ec.stub("keymap--get-keyelt") // Source file: keymap.c
}

func (ec *execContext) keymapParent_autogen(keymap lispObject) (lispObject, error) {
	return ec.stub("keymap-parent") // Source file: keymap.c
}

func (ec *execContext) keymapPrompt_autogen(map_ lispObject) (lispObject, error) {
	return ec.stub("keymap-prompt") // Source file: keymap.c
}

func (ec *execContext) keymapp_autogen(object lispObject) (lispObject, error) {
	return ec.stub("keymapp") // Source file: keymap.c
}

func (ec *execContext) lookupKey_autogen(keymap, key, accept_default lispObject) (lispObject, error) {
	return ec.stub("lookup-key") // Source file: keymap.c
}

func (ec *execContext) makeKeymap_autogen(string lispObject) (lispObject, error) {
	return ec.stub("make-keymap") // Source file: keymap.c
}

func (ec *execContext) makeSparseKeymap_autogen(string lispObject) (lispObject, error) {
	return ec.stub("make-sparse-keymap") // Source file: keymap.c
}

func (ec *execContext) mapKeymap_autogen(function, keymap, sort_first lispObject) (lispObject, error) {
	return ec.stub("map-keymap") // Source file: keymap.c
}

func (ec *execContext) mapKeymapInternal_autogen(function, keymap lispObject) (lispObject, error) {
	return ec.stub("map-keymap-internal") // Source file: keymap.c
}

func (ec *execContext) minorModeKeyBinding_autogen(key, accept_default lispObject) (lispObject, error) {
	return ec.stub("minor-mode-key-binding") // Source file: keymap.c
}

func (ec *execContext) setKeymapParent_autogen(keymap, parent lispObject) (lispObject, error) {
	return ec.stub("set-keymap-parent") // Source file: keymap.c
}

func (ec *execContext) singleKeyDescription_autogen(key, no_angles lispObject) (lispObject, error) {
	return ec.stub("single-key-description") // Source file: keymap.c
}

func (ec *execContext) textCharDescription_autogen(character lispObject) (lispObject, error) {
	return ec.stub("text-char-description") // Source file: keymap.c
}

func (ec *execContext) useGlobalMap_autogen(keymap lispObject) (lispObject, error) {
	return ec.stub("use-global-map") // Source file: keymap.c
}

func (ec *execContext) useLocalMap_autogen(keymap lispObject) (lispObject, error) {
	return ec.stub("use-local-map") // Source file: keymap.c
}

func (ec *execContext) whereIsInternal_autogen(definition, keymap, firstonly, noindirect, no_remap lispObject) (lispObject, error) {
	return ec.stub("where-is-internal") // Source file: keymap.c
}

func (ec *execContext) kqueueAddWatch_autogen(file, flags, callback lispObject) (lispObject, error) {
	return ec.stub("kqueue-add-watch") // Source file: kqueue.c
}

func (ec *execContext) kqueueRmWatch_autogen(watch_descriptor lispObject) (lispObject, error) {
	return ec.stub("kqueue-rm-watch") // Source file: kqueue.c
}

func (ec *execContext) kqueueValidP_autogen(watch_descriptor lispObject) (lispObject, error) {
	return ec.stub("kqueue-valid-p") // Source file: kqueue.c
}

func (ec *execContext) lcmsCam02Ucs_autogen(color1, color2, whitepoint, view lispObject) (lispObject, error) {
	return ec.stub("lcms-cam02-ucs") // Source file: lcms.c
}

func (ec *execContext) lcmsCieDe2000_autogen(color1, color2, kL, kC, kH lispObject) (lispObject, error) {
	return ec.stub("lcms-cie-de2000") // Source file: lcms.c
}

func (ec *execContext) lcmsJabToJch_autogen(color, whitepoint, view lispObject) (lispObject, error) {
	return ec.stub("lcms-jab->jch") // Source file: lcms.c
}

func (ec *execContext) lcmsJchToJab_autogen(color, whitepoint, view lispObject) (lispObject, error) {
	return ec.stub("lcms-jch->jab") // Source file: lcms.c
}

func (ec *execContext) lcmsJchToXyz_autogen(color, whitepoint, view lispObject) (lispObject, error) {
	return ec.stub("lcms-jch->xyz") // Source file: lcms.c
}

func (ec *execContext) lcmsTempToWhitePoint_autogen(temperature lispObject) (lispObject, error) {
	return ec.stub("lcms-temp->white-point") // Source file: lcms.c
}

func (ec *execContext) lcmsXyzToJch_autogen(color, whitepoint, view lispObject) (lispObject, error) {
	return ec.stub("lcms-xyz->jch") // Source file: lcms.c
}

func (ec *execContext) lcms2AvailableP_autogen() (lispObject, error) {
	return ec.stub("lcms2-available-p") // Source file: lcms.c
}

func (ec *execContext) evalBuffer_autogen(buffer, printflag, filename, unibyte, do_allow_print lispObject) (lispObject, error) {
	return ec.stub("eval-buffer") // Source file: lread.c
}

func (ec *execContext) evalRegion_autogen(start, end, printflag, read_function lispObject) (lispObject, error) {
	return ec.stub("eval-region") // Source file: lread.c
}

func (ec *execContext) getFileChar_autogen() (lispObject, error) {
	return ec.stub("get-file-char") // Source file: lread.c
}

func (ec *execContext) getLoadSuffixes_autogen() (lispObject, error) {
	return ec.stub("get-load-suffixes") // Source file: lread.c
}

func (ec *execContext) intern_autogen(string, obarray lispObject) (lispObject, error) {
	return ec.stub("intern") // Source file: lread.c
}

func (ec *execContext) internSoft_autogen(name, obarray lispObject) (lispObject, error) {
	return ec.stub("intern-soft") // Source file: lread.c
}

func (ec *execContext) load_autogen(file, noerror, nomessage, nosuffix, must_suffix lispObject) (lispObject, error) {
	return ec.stub("load") // Source file: lread.c
}

func (ec *execContext) locateFileInternal_autogen(filename, path, suffixes, predicate lispObject) (lispObject, error) {
	return ec.stub("locate-file-internal") // Source file: lread.c
}

func (ec *execContext) lreadSubstituteObjectInSubtree_autogen(object, placeholder, completed lispObject) (lispObject, error) {
	return ec.stub("lread--substitute-object-in-subtree") // Source file: lread.c
}

func (ec *execContext) mapatoms_autogen(function, obarray lispObject) (lispObject, error) {
	return ec.stub("mapatoms") // Source file: lread.c
}

func (ec *execContext) read_autogen(stream lispObject) (lispObject, error) {
	return ec.stub("read") // Source file: lread.c
}

func (ec *execContext) readChar_autogen(prompt, inherit_input_method, seconds lispObject) (lispObject, error) {
	return ec.stub("read-char") // Source file: lread.c
}

func (ec *execContext) readCharExclusive_autogen(prompt, inherit_input_method, seconds lispObject) (lispObject, error) {
	return ec.stub("read-char-exclusive") // Source file: lread.c
}

func (ec *execContext) readEvent_autogen(prompt, inherit_input_method, seconds lispObject) (lispObject, error) {
	return ec.stub("read-event") // Source file: lread.c
}

func (ec *execContext) readFromString_autogen(string, start, end lispObject) (lispObject, error) {
	return ec.stub("read-from-string") // Source file: lread.c
}

func (ec *execContext) readPositioningSymbols_autogen(stream lispObject) (lispObject, error) {
	return ec.stub("read-positioning-symbols") // Source file: lread.c
}

func (ec *execContext) unintern_autogen(name, obarray lispObject) (lispObject, error) {
	return ec.stub("unintern") // Source file: lread.c
}

func (ec *execContext) callLastKbdMacro_autogen(prefix, loopfunc lispObject) (lispObject, error) {
	return ec.stub("call-last-kbd-macro") // Source file: macros.c
}

func (ec *execContext) cancelKbdMacroEvents_autogen() (lispObject, error) {
	return ec.stub("cancel-kbd-macro-events") // Source file: macros.c
}

func (ec *execContext) endKbdMacro_autogen(repeat, loopfunc lispObject) (lispObject, error) {
	return ec.stub("end-kbd-macro") // Source file: macros.c
}

func (ec *execContext) executeKbdMacro_autogen(macro, count, loopfunc lispObject) (lispObject, error) {
	return ec.stub("execute-kbd-macro") // Source file: macros.c
}

func (ec *execContext) startKbdMacro_autogen(append, no_exec lispObject) (lispObject, error) {
	return ec.stub("start-kbd-macro") // Source file: macros.c
}

func (ec *execContext) storeKbdMacroEvent_autogen(event lispObject) (lispObject, error) {
	return ec.stub("store-kbd-macro-event") // Source file: macros.c
}

func (ec *execContext) copyMarker_autogen(marker, type_ lispObject) (lispObject, error) {
	return ec.stub("copy-marker") // Source file: marker.c
}

func (ec *execContext) markerBuffer_autogen(marker lispObject) (lispObject, error) {
	return ec.stub("marker-buffer") // Source file: marker.c
}

func (ec *execContext) markerInsertionType_autogen(marker lispObject) (lispObject, error) {
	return ec.stub("marker-insertion-type") // Source file: marker.c
}

func (ec *execContext) markerPosition_autogen(marker lispObject) (lispObject, error) {
	return ec.stub("marker-position") // Source file: marker.c
}

func (ec *execContext) setMarker_autogen(marker, position, buffer lispObject) (lispObject, error) {
	return ec.stub("set-marker") // Source file: marker.c
}

func (ec *execContext) setMarkerInsertionType_autogen(marker, type_ lispObject) (lispObject, error) {
	return ec.stub("set-marker-insertion-type") // Source file: marker.c
}

func (ec *execContext) menuBarMenuAtXY_autogen(x, y, frame lispObject) (lispObject, error) {
	return ec.stub("menu-bar-menu-at-x-y") // Source file: menu.c
}

func (ec *execContext) xPopupDialog_autogen(position, contents, header lispObject) (lispObject, error) {
	return ec.stub("x-popup-dialog") // Source file: menu.c
}

func (ec *execContext) xPopupMenu_autogen(position, menu lispObject) (lispObject, error) {
	return ec.stub("x-popup-menu") // Source file: menu.c
}

func (ec *execContext) abortMinibuffers_autogen() (lispObject, error) {
	return ec.stub("abort-minibuffers") // Source file: minibuf.c
}

func (ec *execContext) activeMinibufferWindow_autogen() (lispObject, error) {
	return ec.stub("active-minibuffer-window") // Source file: minibuf.c
}

func (ec *execContext) allCompletions_autogen(string, collection, predicate, hide_spaces lispObject) (lispObject, error) {
	return ec.stub("all-completions") // Source file: minibuf.c
}

func (ec *execContext) assocString_autogen(key, list, case_fold lispObject) (lispObject, error) {
	return ec.stub("assoc-string") // Source file: minibuf.c
}

func (ec *execContext) completingRead_autogen(prompt, collection, predicate, require_match, initial_input, hist, def, inherit_input_method lispObject) (lispObject, error) {
	return ec.stub("completing-read") // Source file: minibuf.c
}

func (ec *execContext) innermostMinibufferP_autogen(buffer lispObject) (lispObject, error) {
	return ec.stub("innermost-minibuffer-p") // Source file: minibuf.c
}

func (ec *execContext) internalCompleteBuffer_autogen(string, predicate, flag lispObject) (lispObject, error) {
	return ec.stub("internal-complete-buffer") // Source file: minibuf.c
}

func (ec *execContext) minibufferContents_autogen() (lispObject, error) {
	return ec.stub("minibuffer-contents") // Source file: minibuf.c
}

func (ec *execContext) minibufferContentsNoProperties_autogen() (lispObject, error) {
	return ec.stub("minibuffer-contents-no-properties") // Source file: minibuf.c
}

func (ec *execContext) minibufferDepth_autogen() (lispObject, error) {
	return ec.stub("minibuffer-depth") // Source file: minibuf.c
}

func (ec *execContext) minibufferInnermostCommandLoopP_autogen(buffer lispObject) (lispObject, error) {
	return ec.stub("minibuffer-innermost-command-loop-p") // Source file: minibuf.c
}

func (ec *execContext) minibufferPrompt_autogen() (lispObject, error) {
	return ec.stub("minibuffer-prompt") // Source file: minibuf.c
}

func (ec *execContext) minibufferPromptEnd_autogen() (lispObject, error) {
	return ec.stub("minibuffer-prompt-end") // Source file: minibuf.c
}

func (ec *execContext) minibufferp_autogen(buffer, live lispObject) (lispObject, error) {
	return ec.stub("minibufferp") // Source file: minibuf.c
}

func (ec *execContext) readBuffer_autogen(prompt, def, require_match, predicate lispObject) (lispObject, error) {
	return ec.stub("read-buffer") // Source file: minibuf.c
}

func (ec *execContext) readCommand_autogen(prompt, default_value lispObject) (lispObject, error) {
	return ec.stub("read-command") // Source file: minibuf.c
}

func (ec *execContext) readFromMinibuffer_autogen(prompt, initial_contents, keymap, read, hist, default_value, inherit_input_method lispObject) (lispObject, error) {
	return ec.stub("read-from-minibuffer") // Source file: minibuf.c
}

func (ec *execContext) readFunction_autogen(prompt lispObject) (lispObject, error) {
	return ec.stub("read-function") // Source file: minibuf.c
}

func (ec *execContext) readString_autogen(prompt, initial_input, history, default_value, inherit_input_method lispObject) (lispObject, error) {
	return ec.stub("read-string") // Source file: minibuf.c
}

func (ec *execContext) readVariable_autogen(prompt, default_value lispObject) (lispObject, error) {
	return ec.stub("read-variable") // Source file: minibuf.c
}

func (ec *execContext) setMinibufferWindow_autogen(window lispObject) (lispObject, error) {
	return ec.stub("set-minibuffer-window") // Source file: minibuf.c
}

func (ec *execContext) testCompletion_autogen(string, collection, predicate lispObject) (lispObject, error) {
	return ec.stub("test-completion") // Source file: minibuf.c
}

func (ec *execContext) tryCompletion_autogen(string, collection, predicate lispObject) (lispObject, error) {
	return ec.stub("try-completion") // Source file: minibuf.c
}

func (ec *execContext) dumpEmacsPortable_autogen(filename, track_referrers lispObject) (lispObject, error) {
	return ec.stub("dump-emacs-portable") // Source file: pdumper.c
}

func (ec *execContext) dumpEmacsPortableSortPredicate_autogen(a, b lispObject) (lispObject, error) {
	return ec.stub("dump-emacs-portable--sort-predicate") // Source file: pdumper.c
}

func (ec *execContext) dumpEmacsPortableSortPredicateCopied_autogen(a, b lispObject) (lispObject, error) {
	return ec.stub("dump-emacs-portable--sort-predicate-copied") // Source file: pdumper.c
}

func (ec *execContext) pdumperStats_autogen() (lispObject, error) {
	return ec.stub("pdumper-stats") // Source file: pdumper.c
}

func (ec *execContext) pgtkBackendDisplayClass_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("pgtk-backend-display-class") // Source file: pgtkfns.c
}

func (ec *execContext) pgtkDisplayMonitorAttributesList_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("pgtk-display-monitor-attributes-list") // Source file: pgtkfns.c
}

func (ec *execContext) pgtkFontName_autogen(name lispObject) (lispObject, error) {
	return ec.stub("pgtk-font-name") // Source file: pgtkfns.c
}

func (ec *execContext) pgtkFrameEdges_autogen(frame, type_ lispObject) (lispObject, error) {
	return ec.stub("pgtk-frame-edges") // Source file: pgtkfns.c
}

func (ec *execContext) pgtkFrameGeometry_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("pgtk-frame-geometry") // Source file: pgtkfns.c
}

func (ec *execContext) pgtkFrameRestack_autogen(frame1, frame2, above lispObject) (lispObject, error) {
	return ec.stub("pgtk-frame-restack") // Source file: pgtkfns.c
}

func (ec *execContext) pgtkGetPageSetup_autogen() (lispObject, error) {
	return ec.stub("pgtk-get-page-setup") // Source file: pgtkfns.c
}

func (ec *execContext) pgtkMouseAbsolutePixelPosition_autogen() (lispObject, error) {
	return ec.stub("pgtk-mouse-absolute-pixel-position") // Source file: pgtkfns.c
}

func (ec *execContext) pgtkPageSetupDialog_autogen() (lispObject, error) {
	return ec.stub("pgtk-page-setup-dialog") // Source file: pgtkfns.c
}

func (ec *execContext) pgtkPrintFramesDialog_autogen(frames lispObject) (lispObject, error) {
	return ec.stub("pgtk-print-frames-dialog") // Source file: pgtkfns.c
}

func (ec *execContext) pgtkSetMonitorScaleFactor_autogen(monitor_model, scale_factor lispObject) (lispObject, error) {
	return ec.stub("pgtk-set-monitor-scale-factor") // Source file: pgtkfns.c
}

func (ec *execContext) pgtkSetMouseAbsolutePixelPosition_autogen(x, y lispObject) (lispObject, error) {
	return ec.stub("pgtk-set-mouse-absolute-pixel-position") // Source file: pgtkfns.c
}

func (ec *execContext) pgtkSetResource_autogen(attribute, value lispObject) (lispObject, error) {
	return ec.stub("pgtk-set-resource") // Source file: pgtkfns.c
}

func (ec *execContext) xExportFrames_autogen(frames, type_ lispObject) (lispObject, error) {
	return ec.stub("x-export-frames") // Source file: pgtkfns.c
}

func (ec *execContext) xExportFrames_1_autogen(frames, type_ lispObject) (lispObject, error) {
	return ec.stub("x-export-frames") // Source file: xfns.c
}

func (ec *execContext) xFileDialog_autogen(prompt, dir, default_filename, mustmatch, only_dir_p lispObject) (lispObject, error) {
	return ec.stub("x-file-dialog") // Source file: pgtkfns.c
}

func (ec *execContext) xFileDialog_1_autogen(prompt, dir, default_filename, mustmatch, only_dir_p lispObject) (lispObject, error) {
	return ec.stub("x-file-dialog") // Source file: w32fns.c
}

func (ec *execContext) xFileDialog_2_autogen(prompt, dir, default_filename, mustmatch, only_dir_p lispObject) (lispObject, error) {
	return ec.stub("x-file-dialog") // Source file: xfns.c
}

func (ec *execContext) xFileDialog_3_autogen(prompt, dir, default_filename, mustmatch, only_dir_p lispObject) (lispObject, error) {
	return ec.stub("x-file-dialog") // Source file: xfns.c
}

func (ec *execContext) xGtkDebug_autogen(enable lispObject) (lispObject, error) {
	return ec.stub("x-gtk-debug") // Source file: pgtkfns.c
}

func (ec *execContext) xGtkDebug_1_autogen(enable lispObject) (lispObject, error) {
	return ec.stub("x-gtk-debug") // Source file: xfns.c
}

func (ec *execContext) xServerMaxRequestSize_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-server-max-request-size") // Source file: pgtkfns.c
}

func (ec *execContext) xServerMaxRequestSize_1_autogen(display lispObject) (lispObject, error) {
	return ec.stub("x-server-max-request-size") // Source file: w32fns.c
}

func (ec *execContext) xServerMaxRequestSize_2_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-server-max-request-size") // Source file: xfns.c
}

func (ec *execContext) pgtkUseImContext_autogen(use_p, terminal lispObject) (lispObject, error) {
	return ec.stub("pgtk-use-im-context") // Source file: pgtkim.c
}

func (ec *execContext) xMenuBarOpenInternal_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("x-menu-bar-open-internal") // Source file: pgtkmenu.c
}

func (ec *execContext) xMenuBarOpenInternal_1_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("x-menu-bar-open-internal") // Source file: xmenu.c
}

func (ec *execContext) xMenuBarOpenInternal_2_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("x-menu-bar-open-internal") // Source file: xmenu.c
}

func (ec *execContext) pgtkDisownSelectionInternal_autogen(selection, time_object, terminal lispObject) (lispObject, error) {
	return ec.stub("pgtk-disown-selection-internal") // Source file: pgtkselect.c
}

func (ec *execContext) pgtkDropFinish_autogen(success, timestamp, delete lispObject) (lispObject, error) {
	return ec.stub("pgtk-drop-finish") // Source file: pgtkselect.c
}

func (ec *execContext) pgtkGetSelectionInternal_autogen(selection_symbol, target_type, time_stamp, terminal lispObject) (lispObject, error) {
	return ec.stub("pgtk-get-selection-internal") // Source file: pgtkselect.c
}

func (ec *execContext) pgtkOwnSelectionInternal_autogen(selection, value, frame lispObject) (lispObject, error) {
	return ec.stub("pgtk-own-selection-internal") // Source file: pgtkselect.c
}

func (ec *execContext) pgtkRegisterDndTargets_autogen(frame, targets lispObject) (lispObject, error) {
	return ec.stub("pgtk-register-dnd-targets") // Source file: pgtkselect.c
}

func (ec *execContext) pgtkSelectionExistsP_autogen(selection, terminal lispObject) (lispObject, error) {
	return ec.stub("pgtk-selection-exists-p") // Source file: pgtkselect.c
}

func (ec *execContext) pgtkSelectionOwnerP_autogen(selection, terminal lispObject) (lispObject, error) {
	return ec.stub("pgtk-selection-owner-p") // Source file: pgtkselect.c
}

func (ec *execContext) pgtkUpdateDropStatus_autogen(action, timestamp lispObject) (lispObject, error) {
	return ec.stub("pgtk-update-drop-status") // Source file: pgtkselect.c
}

func (ec *execContext) errorMessageString_autogen(obj lispObject) (lispObject, error) {
	return ec.stub("error-message-string") // Source file: print.c
}

func (ec *execContext) externalDebuggingOutput_autogen(character lispObject) (lispObject, error) {
	return ec.stub("external-debugging-output") // Source file: print.c
}

func (ec *execContext) flushStandardOutput_autogen() (lispObject, error) {
	return ec.stub("flush-standard-output") // Source file: print.c
}

func (ec *execContext) prin1_autogen(object, printcharfun, overrides lispObject) (lispObject, error) {
	return ec.stub("prin1") // Source file: print.c
}

func (ec *execContext) prin1ToString_autogen(object, noescape, overrides lispObject) (lispObject, error) {
	return ec.stub("prin1-to-string") // Source file: print.c
}

func (ec *execContext) princ_autogen(object, printcharfun lispObject) (lispObject, error) {
	return ec.stub("princ") // Source file: print.c
}

func (ec *execContext) print_autogen(object, printcharfun lispObject) (lispObject, error) {
	return ec.stub("print") // Source file: print.c
}

func (ec *execContext) printPreprocess_autogen(object lispObject) (lispObject, error) {
	return ec.stub("print--preprocess") // Source file: print.c
}

func (ec *execContext) redirectDebuggingOutput_autogen(file, append lispObject) (lispObject, error) {
	return ec.stub("redirect-debugging-output") // Source file: print.c
}

func (ec *execContext) terpri_autogen(printcharfun, ensure lispObject) (lispObject, error) {
	return ec.stub("terpri") // Source file: print.c
}

func (ec *execContext) writeChar_autogen(character, printcharfun lispObject) (lispObject, error) {
	return ec.stub("write-char") // Source file: print.c
}

func (ec *execContext) acceptProcessOutput_autogen(process, seconds, millisec, just_this_one lispObject) (lispObject, error) {
	return ec.stub("accept-process-output") // Source file: process.c
}

func (ec *execContext) continueProcess_autogen(process, current_group lispObject) (lispObject, error) {
	return ec.stub("continue-process") // Source file: process.c
}

func (ec *execContext) deleteProcess_autogen(process lispObject) (lispObject, error) {
	return ec.stub("delete-process") // Source file: process.c
}

func (ec *execContext) formatNetworkAddress_autogen(address, omit_port lispObject) (lispObject, error) {
	return ec.stub("format-network-address") // Source file: process.c
}

func (ec *execContext) getBufferProcess_autogen(buffer lispObject) (lispObject, error) {
	return ec.stub("get-buffer-process") // Source file: process.c
}

func (ec *execContext) getProcess_autogen(name lispObject) (lispObject, error) {
	return ec.stub("get-process") // Source file: process.c
}

func (ec *execContext) internalDefaultInterruptProcess_autogen(process, current_group lispObject) (lispObject, error) {
	return ec.stub("internal-default-interrupt-process") // Source file: process.c
}

func (ec *execContext) internalDefaultProcessFilter_autogen(proc, text lispObject) (lispObject, error) {
	return ec.stub("internal-default-process-filter") // Source file: process.c
}

func (ec *execContext) internalDefaultProcessSentinel_autogen(proc, msg lispObject) (lispObject, error) {
	return ec.stub("internal-default-process-sentinel") // Source file: process.c
}

func (ec *execContext) internalDefaultSignalProcess_autogen(process, sigcode, remote lispObject) (lispObject, error) {
	return ec.stub("internal-default-signal-process") // Source file: process.c
}

func (ec *execContext) interruptProcess_autogen(process, current_group lispObject) (lispObject, error) {
	return ec.stub("interrupt-process") // Source file: process.c
}

func (ec *execContext) killProcess_autogen(process, current_group lispObject) (lispObject, error) {
	return ec.stub("kill-process") // Source file: process.c
}

func (ec *execContext) listSystemProcesses_autogen() (lispObject, error) {
	return ec.stub("list-system-processes") // Source file: process.c
}

func (ec *execContext) makeNetworkProcess_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("make-network-process") // Source file: process.c
}

func (ec *execContext) makePipeProcess_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("make-pipe-process") // Source file: process.c
}

func (ec *execContext) makeProcess_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("make-process") // Source file: process.c
}

func (ec *execContext) makeSerialProcess_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("make-serial-process") // Source file: process.c
}

func (ec *execContext) networkInterfaceInfo_autogen(ifname lispObject) (lispObject, error) {
	return ec.stub("network-interface-info") // Source file: process.c
}

func (ec *execContext) networkInterfaceList_autogen(full, family lispObject) (lispObject, error) {
	return ec.stub("network-interface-list") // Source file: process.c
}

func (ec *execContext) networkLookupAddressInfo_autogen(name, family, hint lispObject) (lispObject, error) {
	return ec.stub("network-lookup-address-info") // Source file: process.c
}

func (ec *execContext) numProcessors_autogen(query lispObject) (lispObject, error) {
	return ec.stub("num-processors") // Source file: process.c
}

func (ec *execContext) processAttributes_autogen(pid lispObject) (lispObject, error) {
	return ec.stub("process-attributes") // Source file: process.c
}

func (ec *execContext) processBuffer_autogen(process lispObject) (lispObject, error) {
	return ec.stub("process-buffer") // Source file: process.c
}

func (ec *execContext) processCodingSystem_autogen(process lispObject) (lispObject, error) {
	return ec.stub("process-coding-system") // Source file: process.c
}

func (ec *execContext) processCommand_autogen(process lispObject) (lispObject, error) {
	return ec.stub("process-command") // Source file: process.c
}

func (ec *execContext) processConnection_autogen(process lispObject) (lispObject, error) {
	return ec.stub("process-connection") // Source file: process.c
}

func (ec *execContext) processContact_autogen(process, key, no_block lispObject) (lispObject, error) {
	return ec.stub("process-contact") // Source file: process.c
}

func (ec *execContext) processDatagramAddress_autogen(process lispObject) (lispObject, error) {
	return ec.stub("process-datagram-address") // Source file: process.c
}

func (ec *execContext) processExitStatus_autogen(process lispObject) (lispObject, error) {
	return ec.stub("process-exit-status") // Source file: process.c
}

func (ec *execContext) processFilter_autogen(process lispObject) (lispObject, error) {
	return ec.stub("process-filter") // Source file: process.c
}

func (ec *execContext) processId_autogen(process lispObject) (lispObject, error) {
	return ec.stub("process-id") // Source file: process.c
}

func (ec *execContext) processInheritCodingSystemFlag_autogen(process lispObject) (lispObject, error) {
	return ec.stub("process-inherit-coding-system-flag") // Source file: process.c
}

func (ec *execContext) processList_autogen() (lispObject, error) {
	return ec.stub("process-list") // Source file: process.c
}

func (ec *execContext) processMark_autogen(process lispObject) (lispObject, error) {
	return ec.stub("process-mark") // Source file: process.c
}

func (ec *execContext) processName_autogen(process lispObject) (lispObject, error) {
	return ec.stub("process-name") // Source file: process.c
}

func (ec *execContext) processPlist_autogen(process lispObject) (lispObject, error) {
	return ec.stub("process-plist") // Source file: process.c
}

func (ec *execContext) processQueryOnExitFlag_autogen(process lispObject) (lispObject, error) {
	return ec.stub("process-query-on-exit-flag") // Source file: process.c
}

func (ec *execContext) processRunningChildP_autogen(process lispObject) (lispObject, error) {
	return ec.stub("process-running-child-p") // Source file: process.c
}

func (ec *execContext) processSendEof_autogen(process lispObject) (lispObject, error) {
	return ec.stub("process-send-eof") // Source file: process.c
}

func (ec *execContext) processSendRegion_autogen(process, start, end lispObject) (lispObject, error) {
	return ec.stub("process-send-region") // Source file: process.c
}

func (ec *execContext) processSendString_autogen(process, string lispObject) (lispObject, error) {
	return ec.stub("process-send-string") // Source file: process.c
}

func (ec *execContext) processSentinel_autogen(process lispObject) (lispObject, error) {
	return ec.stub("process-sentinel") // Source file: process.c
}

func (ec *execContext) processStatus_autogen(process lispObject) (lispObject, error) {
	return ec.stub("process-status") // Source file: process.c
}

func (ec *execContext) processThread_autogen(process lispObject) (lispObject, error) {
	return ec.stub("process-thread") // Source file: process.c
}

func (ec *execContext) processTtyName_autogen(process, stream lispObject) (lispObject, error) {
	return ec.stub("process-tty-name") // Source file: process.c
}

func (ec *execContext) processType_autogen(process lispObject) (lispObject, error) {
	return ec.stub("process-type") // Source file: process.c
}

func (ec *execContext) processp_autogen(object lispObject) (lispObject, error) {
	return ec.stub("processp") // Source file: process.c
}

func (ec *execContext) quitProcess_autogen(process, current_group lispObject) (lispObject, error) {
	return ec.stub("quit-process") // Source file: process.c
}

func (ec *execContext) serialProcessConfigure_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("serial-process-configure") // Source file: process.c
}

func (ec *execContext) setNetworkProcessOption_autogen(process, option, value, no_error lispObject) (lispObject, error) {
	return ec.stub("set-network-process-option") // Source file: process.c
}

func (ec *execContext) setProcessBuffer_autogen(process, buffer lispObject) (lispObject, error) {
	return ec.stub("set-process-buffer") // Source file: process.c
}

func (ec *execContext) setProcessCodingSystem_autogen(process, decoding, encoding lispObject) (lispObject, error) {
	return ec.stub("set-process-coding-system") // Source file: process.c
}

func (ec *execContext) setProcessDatagramAddress_autogen(process, address lispObject) (lispObject, error) {
	return ec.stub("set-process-datagram-address") // Source file: process.c
}

func (ec *execContext) setProcessFilter_autogen(process, filter lispObject) (lispObject, error) {
	return ec.stub("set-process-filter") // Source file: process.c
}

func (ec *execContext) setProcessInheritCodingSystemFlag_autogen(process, flag lispObject) (lispObject, error) {
	return ec.stub("set-process-inherit-coding-system-flag") // Source file: process.c
}

func (ec *execContext) setProcessPlist_autogen(process, plist lispObject) (lispObject, error) {
	return ec.stub("set-process-plist") // Source file: process.c
}

func (ec *execContext) setProcessQueryOnExitFlag_autogen(process, flag lispObject) (lispObject, error) {
	return ec.stub("set-process-query-on-exit-flag") // Source file: process.c
}

func (ec *execContext) setProcessSentinel_autogen(process, sentinel lispObject) (lispObject, error) {
	return ec.stub("set-process-sentinel") // Source file: process.c
}

func (ec *execContext) setProcessThread_autogen(process, thread lispObject) (lispObject, error) {
	return ec.stub("set-process-thread") // Source file: process.c
}

func (ec *execContext) setProcessWindowSize_autogen(process, height, width lispObject) (lispObject, error) {
	return ec.stub("set-process-window-size") // Source file: process.c
}

func (ec *execContext) signalNames_autogen() (lispObject, error) {
	return ec.stub("signal-names") // Source file: process.c
}

func (ec *execContext) signalProcess_autogen(process, sigcode, remote lispObject) (lispObject, error) {
	return ec.stub("signal-process") // Source file: process.c
}

func (ec *execContext) stopProcess_autogen(process, current_group lispObject) (lispObject, error) {
	return ec.stub("stop-process") // Source file: process.c
}

func (ec *execContext) waitingForUserInputP_autogen() (lispObject, error) {
	return ec.stub("waiting-for-user-input-p") // Source file: process.c
}

func (ec *execContext) functionEqual_autogen(f1, f2 lispObject) (lispObject, error) {
	return ec.stub("function-equal") // Source file: profiler.c
}

func (ec *execContext) profilerCpuLog_autogen() (lispObject, error) {
	return ec.stub("profiler-cpu-log") // Source file: profiler.c
}

func (ec *execContext) profilerCpuRunningP_autogen() (lispObject, error) {
	return ec.stub("profiler-cpu-running-p") // Source file: profiler.c
}

func (ec *execContext) profilerCpuStart_autogen(sampling_interval lispObject) (lispObject, error) {
	return ec.stub("profiler-cpu-start") // Source file: profiler.c
}

func (ec *execContext) profilerCpuStop_autogen() (lispObject, error) {
	return ec.stub("profiler-cpu-stop") // Source file: profiler.c
}

func (ec *execContext) profilerMemoryLog_autogen() (lispObject, error) {
	return ec.stub("profiler-memory-log") // Source file: profiler.c
}

func (ec *execContext) profilerMemoryRunningP_autogen() (lispObject, error) {
	return ec.stub("profiler-memory-running-p") // Source file: profiler.c
}

func (ec *execContext) profilerMemoryStart_autogen() (lispObject, error) {
	return ec.stub("profiler-memory-start") // Source file: profiler.c
}

func (ec *execContext) profilerMemoryStop_autogen() (lispObject, error) {
	return ec.stub("profiler-memory-stop") // Source file: profiler.c
}

func (ec *execContext) lookingAt_autogen(regexp, inhibit_modify lispObject) (lispObject, error) {
	return ec.stub("looking-at") // Source file: search.c
}

func (ec *execContext) matchBeginning_autogen(subexp lispObject) (lispObject, error) {
	return ec.stub("match-beginning") // Source file: search.c
}

func (ec *execContext) matchData_autogen(integers, reuse, reseat lispObject) (lispObject, error) {
	return ec.stub("match-data") // Source file: search.c
}

func (ec *execContext) matchDataTranslate_autogen(n lispObject) (lispObject, error) {
	return ec.stub("match-data--translate") // Source file: search.c
}

func (ec *execContext) matchEnd_autogen(subexp lispObject) (lispObject, error) {
	return ec.stub("match-end") // Source file: search.c
}

func (ec *execContext) newlineCacheCheck_autogen(buffer lispObject) (lispObject, error) {
	return ec.stub("newline-cache-check") // Source file: search.c
}

func (ec *execContext) posixLookingAt_autogen(regexp, inhibit_modify lispObject) (lispObject, error) {
	return ec.stub("posix-looking-at") // Source file: search.c
}

func (ec *execContext) posixSearchBackward_autogen(regexp, bound, noerror, count lispObject) (lispObject, error) {
	return ec.stub("posix-search-backward") // Source file: search.c
}

func (ec *execContext) posixSearchForward_autogen(regexp, bound, noerror, count lispObject) (lispObject, error) {
	return ec.stub("posix-search-forward") // Source file: search.c
}

func (ec *execContext) posixStringMatch_autogen(regexp, string, start, inhibit_modify lispObject) (lispObject, error) {
	return ec.stub("posix-string-match") // Source file: search.c
}

func (ec *execContext) reSearchBackward_autogen(regexp, bound, noerror, count lispObject) (lispObject, error) {
	return ec.stub("re-search-backward") // Source file: search.c
}

func (ec *execContext) reSearchForward_autogen(regexp, bound, noerror, count lispObject) (lispObject, error) {
	return ec.stub("re-search-forward") // Source file: search.c
}

func (ec *execContext) regexpQuote_autogen(string lispObject) (lispObject, error) {
	return ec.stub("regexp-quote") // Source file: search.c
}

func (ec *execContext) replaceMatch_autogen(newtext, fixedcase, literal, string, subexp lispObject) (lispObject, error) {
	return ec.stub("replace-match") // Source file: search.c
}

func (ec *execContext) searchBackward_autogen(string, bound, noerror, count lispObject) (lispObject, error) {
	return ec.stub("search-backward") // Source file: search.c
}

func (ec *execContext) searchForward_autogen(string, bound, noerror, count lispObject) (lispObject, error) {
	return ec.stub("search-forward") // Source file: search.c
}

func (ec *execContext) setMatchData_autogen(list, reseat lispObject) (lispObject, error) {
	return ec.stub("set-match-data") // Source file: search.c
}

func (ec *execContext) stringMatch_autogen(regexp, string, start, inhibit_modify lispObject) (lispObject, error) {
	return ec.stub("string-match") // Source file: search.c
}

func (ec *execContext) playSoundInternal_autogen(sound lispObject) (lispObject, error) {
	return ec.stub("play-sound-internal") // Source file: sound.c
}

func (ec *execContext) sqliteAvailableP_autogen() (lispObject, error) {
	return ec.stub("sqlite-available-p") // Source file: sqlite.c
}

func (ec *execContext) sqliteClose_autogen(db lispObject) (lispObject, error) {
	return ec.stub("sqlite-close") // Source file: sqlite.c
}

func (ec *execContext) sqliteColumns_autogen(set lispObject) (lispObject, error) {
	return ec.stub("sqlite-columns") // Source file: sqlite.c
}

func (ec *execContext) sqliteCommit_autogen(db lispObject) (lispObject, error) {
	return ec.stub("sqlite-commit") // Source file: sqlite.c
}

func (ec *execContext) sqliteExecute_autogen(db, query, values lispObject) (lispObject, error) {
	return ec.stub("sqlite-execute") // Source file: sqlite.c
}

func (ec *execContext) sqliteFinalize_autogen(set lispObject) (lispObject, error) {
	return ec.stub("sqlite-finalize") // Source file: sqlite.c
}

func (ec *execContext) sqliteLoadExtension_autogen(db, module lispObject) (lispObject, error) {
	return ec.stub("sqlite-load-extension") // Source file: sqlite.c
}

func (ec *execContext) sqliteMoreP_autogen(set lispObject) (lispObject, error) {
	return ec.stub("sqlite-more-p") // Source file: sqlite.c
}

func (ec *execContext) sqliteNext_autogen(set lispObject) (lispObject, error) {
	return ec.stub("sqlite-next") // Source file: sqlite.c
}

func (ec *execContext) sqliteOpen_autogen(file lispObject) (lispObject, error) {
	return ec.stub("sqlite-open") // Source file: sqlite.c
}

func (ec *execContext) sqlitePragma_autogen(db, pragma lispObject) (lispObject, error) {
	return ec.stub("sqlite-pragma") // Source file: sqlite.c
}

func (ec *execContext) sqliteRollback_autogen(db lispObject) (lispObject, error) {
	return ec.stub("sqlite-rollback") // Source file: sqlite.c
}

func (ec *execContext) sqliteSelect_autogen(db, query, values, return_type lispObject) (lispObject, error) {
	return ec.stub("sqlite-select") // Source file: sqlite.c
}

func (ec *execContext) sqliteTransaction_autogen(db lispObject) (lispObject, error) {
	return ec.stub("sqlite-transaction") // Source file: sqlite.c
}

func (ec *execContext) sqliteVersion_autogen() (lispObject, error) {
	return ec.stub("sqlite-version") // Source file: sqlite.c
}

func (ec *execContext) sqlitep_autogen(object lispObject) (lispObject, error) {
	return ec.stub("sqlitep") // Source file: sqlite.c
}

func (ec *execContext) backwardPrefixChars_autogen() (lispObject, error) {
	return ec.stub("backward-prefix-chars") // Source file: syntax.c
}

func (ec *execContext) charSyntax_autogen(character lispObject) (lispObject, error) {
	return ec.stub("char-syntax") // Source file: syntax.c
}

func (ec *execContext) copySyntaxTable_autogen(table lispObject) (lispObject, error) {
	return ec.stub("copy-syntax-table") // Source file: syntax.c
}

func (ec *execContext) forwardComment_autogen(count lispObject) (lispObject, error) {
	return ec.stub("forward-comment") // Source file: syntax.c
}

func (ec *execContext) forwardWord_autogen(arg lispObject) (lispObject, error) {
	return ec.stub("forward-word") // Source file: syntax.c
}

func (ec *execContext) internalDescribeSyntaxValue_autogen(syntax lispObject) (lispObject, error) {
	return ec.stub("internal-describe-syntax-value") // Source file: syntax.c
}

func (ec *execContext) matchingParen_autogen(character lispObject) (lispObject, error) {
	return ec.stub("matching-paren") // Source file: syntax.c
}

func (ec *execContext) modifySyntaxEntry_autogen(c, newentry, syntax_table lispObject) (lispObject, error) {
	return ec.stub("modify-syntax-entry") // Source file: syntax.c
}

func (ec *execContext) parsePartialSexp_autogen(from, to, targetdepth, stopbefore, oldstate, commentstop lispObject) (lispObject, error) {
	return ec.stub("parse-partial-sexp") // Source file: syntax.c
}

func (ec *execContext) scanLists_autogen(from, count, depth lispObject) (lispObject, error) {
	return ec.stub("scan-lists") // Source file: syntax.c
}

func (ec *execContext) scanSexps_autogen(from, count lispObject) (lispObject, error) {
	return ec.stub("scan-sexps") // Source file: syntax.c
}

func (ec *execContext) setSyntaxTable_autogen(table lispObject) (lispObject, error) {
	return ec.stub("set-syntax-table") // Source file: syntax.c
}

func (ec *execContext) skipCharsBackward_autogen(string, lim lispObject) (lispObject, error) {
	return ec.stub("skip-chars-backward") // Source file: syntax.c
}

func (ec *execContext) skipCharsForward_autogen(string, lim lispObject) (lispObject, error) {
	return ec.stub("skip-chars-forward") // Source file: syntax.c
}

func (ec *execContext) skipSyntaxBackward_autogen(syntax, lim lispObject) (lispObject, error) {
	return ec.stub("skip-syntax-backward") // Source file: syntax.c
}

func (ec *execContext) skipSyntaxForward_autogen(syntax, lim lispObject) (lispObject, error) {
	return ec.stub("skip-syntax-forward") // Source file: syntax.c
}

func (ec *execContext) standardSyntaxTable_autogen() (lispObject, error) {
	return ec.stub("standard-syntax-table") // Source file: syntax.c
}

func (ec *execContext) stringToSyntax_autogen(string lispObject) (lispObject, error) {
	return ec.stub("string-to-syntax") // Source file: syntax.c
}

func (ec *execContext) syntaxClassToChar_autogen(syntax lispObject) (lispObject, error) {
	return ec.stub("syntax-class-to-char") // Source file: syntax.c
}

func (ec *execContext) syntaxTable_autogen() (lispObject, error) {
	return ec.stub("syntax-table") // Source file: syntax.c
}

func (ec *execContext) syntaxTableP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("syntax-table-p") // Source file: syntax.c
}

func (ec *execContext) getInternalRunTime_autogen() (lispObject, error) {
	return ec.stub("get-internal-run-time") // Source file: sysdep.c
}

func (ec *execContext) controllingTtyP_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("controlling-tty-p") // Source file: term.c
}

func (ec *execContext) gpmMouseStart_autogen() (lispObject, error) {
	return ec.stub("gpm-mouse-start") // Source file: term.c
}

func (ec *execContext) gpmMouseStop_autogen() (lispObject, error) {
	return ec.stub("gpm-mouse-stop") // Source file: term.c
}

func (ec *execContext) resumeTty_autogen(tty lispObject) (lispObject, error) {
	return ec.stub("resume-tty") // Source file: term.c
}

func (ec *execContext) suspendTty_autogen(tty lispObject) (lispObject, error) {
	return ec.stub("suspend-tty") // Source file: term.c
}

func (ec *execContext) ttyOutputBufferSize_autogen(tty lispObject) (lispObject, error) {
	return ec.stub("tty--output-buffer-size") // Source file: term.c
}

func (ec *execContext) ttySetOutputBufferSize_autogen(size, tty lispObject) (lispObject, error) {
	return ec.stub("tty--set-output-buffer-size") // Source file: term.c
}

func (ec *execContext) ttyDisplayColorCells_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("tty-display-color-cells") // Source file: term.c
}

func (ec *execContext) ttyDisplayColorP_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("tty-display-color-p") // Source file: term.c
}

func (ec *execContext) ttyNoUnderline_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("tty-no-underline") // Source file: term.c
}

func (ec *execContext) ttyTopFrame_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("tty-top-frame") // Source file: term.c
}

func (ec *execContext) ttyType_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("tty-type") // Source file: term.c
}

func (ec *execContext) deleteTerminal_autogen(terminal, force lispObject) (lispObject, error) {
	return ec.stub("delete-terminal") // Source file: terminal.c
}

func (ec *execContext) frameTerminal_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("frame-terminal") // Source file: terminal.c
}

func (ec *execContext) setTerminalParameter_autogen(terminal, parameter, value lispObject) (lispObject, error) {
	return ec.stub("set-terminal-parameter") // Source file: terminal.c
}

func (ec *execContext) terminalList_autogen() (lispObject, error) {
	return ec.stub("terminal-list") // Source file: terminal.c
}

func (ec *execContext) terminalLiveP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("terminal-live-p") // Source file: terminal.c
}

func (ec *execContext) terminalName_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("terminal-name") // Source file: terminal.c
}

func (ec *execContext) terminalParameter_autogen(terminal, parameter lispObject) (lispObject, error) {
	return ec.stub("terminal-parameter") // Source file: terminal.c
}

func (ec *execContext) terminalParameters_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("terminal-parameters") // Source file: terminal.c
}

func (ec *execContext) addFaceTextProperty_autogen(start, end, face, append, object lispObject) (lispObject, error) {
	return ec.stub("add-face-text-property") // Source file: textprop.c
}

func (ec *execContext) addTextProperties_autogen(start, end, properties, object lispObject) (lispObject, error) {
	return ec.stub("add-text-properties") // Source file: textprop.c
}

func (ec *execContext) getCharProperty_autogen(position, prop, object lispObject) (lispObject, error) {
	return ec.stub("get-char-property") // Source file: textprop.c
}

func (ec *execContext) getCharPropertyAndOverlay_autogen(position, prop, object lispObject) (lispObject, error) {
	return ec.stub("get-char-property-and-overlay") // Source file: textprop.c
}

func (ec *execContext) getTextProperty_autogen(position, prop, object lispObject) (lispObject, error) {
	return ec.stub("get-text-property") // Source file: textprop.c
}

func (ec *execContext) nextCharPropertyChange_autogen(position, limit lispObject) (lispObject, error) {
	return ec.stub("next-char-property-change") // Source file: textprop.c
}

func (ec *execContext) nextPropertyChange_autogen(position, object, limit lispObject) (lispObject, error) {
	return ec.stub("next-property-change") // Source file: textprop.c
}

func (ec *execContext) nextSingleCharPropertyChange_autogen(position, prop, object, limit lispObject) (lispObject, error) {
	return ec.stub("next-single-char-property-change") // Source file: textprop.c
}

func (ec *execContext) nextSinglePropertyChange_autogen(position, prop, object, limit lispObject) (lispObject, error) {
	return ec.stub("next-single-property-change") // Source file: textprop.c
}

func (ec *execContext) previousCharPropertyChange_autogen(position, limit lispObject) (lispObject, error) {
	return ec.stub("previous-char-property-change") // Source file: textprop.c
}

func (ec *execContext) previousPropertyChange_autogen(position, object, limit lispObject) (lispObject, error) {
	return ec.stub("previous-property-change") // Source file: textprop.c
}

func (ec *execContext) previousSingleCharPropertyChange_autogen(position, prop, object, limit lispObject) (lispObject, error) {
	return ec.stub("previous-single-char-property-change") // Source file: textprop.c
}

func (ec *execContext) previousSinglePropertyChange_autogen(position, prop, object, limit lispObject) (lispObject, error) {
	return ec.stub("previous-single-property-change") // Source file: textprop.c
}

func (ec *execContext) putTextProperty_autogen(start, end, property, value, object lispObject) (lispObject, error) {
	return ec.stub("put-text-property") // Source file: textprop.c
}

func (ec *execContext) removeListOfTextProperties_autogen(start, end, list_of_properties, object lispObject) (lispObject, error) {
	return ec.stub("remove-list-of-text-properties") // Source file: textprop.c
}

func (ec *execContext) removeTextProperties_autogen(start, end, properties, object lispObject) (lispObject, error) {
	return ec.stub("remove-text-properties") // Source file: textprop.c
}

func (ec *execContext) setTextProperties_autogen(start, end, properties, object lispObject) (lispObject, error) {
	return ec.stub("set-text-properties") // Source file: textprop.c
}

func (ec *execContext) textPropertiesAt_autogen(position, object lispObject) (lispObject, error) {
	return ec.stub("text-properties-at") // Source file: textprop.c
}

func (ec *execContext) textPropertyAny_autogen(start, end, property, value, object lispObject) (lispObject, error) {
	return ec.stub("text-property-any") // Source file: textprop.c
}

func (ec *execContext) textPropertyNotAll_autogen(start, end, property, value, object lispObject) (lispObject, error) {
	return ec.stub("text-property-not-all") // Source file: textprop.c
}

func (ec *execContext) allThreads_autogen() (lispObject, error) {
	return ec.stub("all-threads") // Source file: thread.c
}

func (ec *execContext) conditionMutex_autogen(cond lispObject) (lispObject, error) {
	return ec.stub("condition-mutex") // Source file: thread.c
}

func (ec *execContext) conditionName_autogen(cond lispObject) (lispObject, error) {
	return ec.stub("condition-name") // Source file: thread.c
}

func (ec *execContext) conditionNotify_autogen(cond, all lispObject) (lispObject, error) {
	return ec.stub("condition-notify") // Source file: thread.c
}

func (ec *execContext) conditionWait_autogen(cond lispObject) (lispObject, error) {
	return ec.stub("condition-wait") // Source file: thread.c
}

func (ec *execContext) currentThread_autogen() (lispObject, error) {
	return ec.stub("current-thread") // Source file: thread.c
}

func (ec *execContext) makeConditionVariable_autogen(mutex, name lispObject) (lispObject, error) {
	return ec.stub("make-condition-variable") // Source file: thread.c
}

func (ec *execContext) makeMutex_autogen(name lispObject) (lispObject, error) {
	return ec.stub("make-mutex") // Source file: thread.c
}

func (ec *execContext) makeThread_autogen(function, name lispObject) (lispObject, error) {
	return ec.stub("make-thread") // Source file: thread.c
}

func (ec *execContext) mutexLock_autogen(mutex lispObject) (lispObject, error) {
	return ec.stub("mutex-lock") // Source file: thread.c
}

func (ec *execContext) mutexName_autogen(mutex lispObject) (lispObject, error) {
	return ec.stub("mutex-name") // Source file: thread.c
}

func (ec *execContext) mutexUnlock_autogen(mutex lispObject) (lispObject, error) {
	return ec.stub("mutex-unlock") // Source file: thread.c
}

func (ec *execContext) threadBlocker_autogen(thread lispObject) (lispObject, error) {
	return ec.stub("thread--blocker") // Source file: thread.c
}

func (ec *execContext) threadJoin_autogen(thread lispObject) (lispObject, error) {
	return ec.stub("thread-join") // Source file: thread.c
}

func (ec *execContext) threadLastError_autogen(cleanup lispObject) (lispObject, error) {
	return ec.stub("thread-last-error") // Source file: thread.c
}

func (ec *execContext) threadLiveP_autogen(thread lispObject) (lispObject, error) {
	return ec.stub("thread-live-p") // Source file: thread.c
}

func (ec *execContext) threadName_autogen(thread lispObject) (lispObject, error) {
	return ec.stub("thread-name") // Source file: thread.c
}

func (ec *execContext) threadSignal_autogen(thread, error_symbol, data lispObject) (lispObject, error) {
	return ec.stub("thread-signal") // Source file: thread.c
}

func (ec *execContext) threadYield_autogen() (lispObject, error) {
	return ec.stub("thread-yield") // Source file: thread.c
}

func (ec *execContext) currentCpuTime_autogen() (lispObject, error) {
	return ec.stub("current-cpu-time") // Source file: timefns.c
}

func (ec *execContext) currentTime_autogen() (lispObject, error) {
	return ec.stub("current-time") // Source file: timefns.c
}

func (ec *execContext) currentTimeString_autogen(specified_time, zone lispObject) (lispObject, error) {
	return ec.stub("current-time-string") // Source file: timefns.c
}

func (ec *execContext) currentTimeZone_autogen(specified_time, zone lispObject) (lispObject, error) {
	return ec.stub("current-time-zone") // Source file: timefns.c
}

func (ec *execContext) decodeTime_autogen(specified_time, zone, form lispObject) (lispObject, error) {
	return ec.stub("decode-time") // Source file: timefns.c
}

func (ec *execContext) encodeTime_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("encode-time") // Source file: timefns.c
}

func (ec *execContext) floatTime_autogen(specified_time lispObject) (lispObject, error) {
	return ec.stub("float-time") // Source file: timefns.c
}

func (ec *execContext) formatTimeString_autogen(format_string, timeval, zone lispObject) (lispObject, error) {
	return ec.stub("format-time-string") // Source file: timefns.c
}

func (ec *execContext) setTimeZoneRule_autogen(tz lispObject) (lispObject, error) {
	return ec.stub("set-time-zone-rule") // Source file: timefns.c
}

func (ec *execContext) timeAdd_autogen(a, b lispObject) (lispObject, error) {
	return ec.stub("time-add") // Source file: timefns.c
}

func (ec *execContext) timeConvert_autogen(time, form lispObject) (lispObject, error) {
	return ec.stub("time-convert") // Source file: timefns.c
}

func (ec *execContext) timeEqualP_autogen(a, b lispObject) (lispObject, error) {
	return ec.stub("time-equal-p") // Source file: timefns.c
}

func (ec *execContext) timeLessP_autogen(a, b lispObject) (lispObject, error) {
	return ec.stub("time-less-p") // Source file: timefns.c
}

func (ec *execContext) timeSubtract_autogen(a, b lispObject) (lispObject, error) {
	return ec.stub("time-subtract") // Source file: timefns.c
}

func (ec *execContext) treesitAvailableP_autogen() (lispObject, error) {
	return ec.stub("treesit-available-p") // Source file: treesit.c
}

func (ec *execContext) treesitCompiledQueryP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("treesit-compiled-query-p") // Source file: treesit.c
}

func (ec *execContext) treesitInduceSparseTree_autogen(root, predicate, process_fn, depth lispObject) (lispObject, error) {
	return ec.stub("treesit-induce-sparse-tree") // Source file: treesit.c
}

func (ec *execContext) treesitLanguageAbiVersion_autogen(language lispObject) (lispObject, error) {
	return ec.stub("treesit-language-abi-version") // Source file: treesit.c
}

func (ec *execContext) treesitLanguageAvailableP_autogen(language, detail lispObject) (lispObject, error) {
	return ec.stub("treesit-language-available-p") // Source file: treesit.c
}

func (ec *execContext) treesitLibraryAbiVersion_autogen(min_compatible lispObject) (lispObject, error) {
	return ec.stub("treesit-library-abi-version") // Source file: treesit.c
}

func (ec *execContext) treesitNodeCheck_autogen(node, property lispObject) (lispObject, error) {
	return ec.stub("treesit-node-check") // Source file: treesit.c
}

func (ec *execContext) treesitNodeChild_autogen(node, n, named lispObject) (lispObject, error) {
	return ec.stub("treesit-node-child") // Source file: treesit.c
}

func (ec *execContext) treesitNodeChildByFieldName_autogen(node, field_name lispObject) (lispObject, error) {
	return ec.stub("treesit-node-child-by-field-name") // Source file: treesit.c
}

func (ec *execContext) treesitNodeChildCount_autogen(node, named lispObject) (lispObject, error) {
	return ec.stub("treesit-node-child-count") // Source file: treesit.c
}

func (ec *execContext) treesitNodeDescendantForRange_autogen(node, beg, end, named lispObject) (lispObject, error) {
	return ec.stub("treesit-node-descendant-for-range") // Source file: treesit.c
}

func (ec *execContext) treesitNodeEnd_autogen(node lispObject) (lispObject, error) {
	return ec.stub("treesit-node-end") // Source file: treesit.c
}

func (ec *execContext) treesitNodeEq_autogen(node1, node2 lispObject) (lispObject, error) {
	return ec.stub("treesit-node-eq") // Source file: treesit.c
}

func (ec *execContext) treesitNodeFieldNameForChild_autogen(node, n lispObject) (lispObject, error) {
	return ec.stub("treesit-node-field-name-for-child") // Source file: treesit.c
}

func (ec *execContext) treesitNodeFirstChildForPos_autogen(node, pos, named lispObject) (lispObject, error) {
	return ec.stub("treesit-node-first-child-for-pos") // Source file: treesit.c
}

func (ec *execContext) treesitNodeNextSibling_autogen(node, named lispObject) (lispObject, error) {
	return ec.stub("treesit-node-next-sibling") // Source file: treesit.c
}

func (ec *execContext) treesitNodeP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("treesit-node-p") // Source file: treesit.c
}

func (ec *execContext) treesitNodeParent_autogen(node lispObject) (lispObject, error) {
	return ec.stub("treesit-node-parent") // Source file: treesit.c
}

func (ec *execContext) treesitNodeParser_autogen(node lispObject) (lispObject, error) {
	return ec.stub("treesit-node-parser") // Source file: treesit.c
}

func (ec *execContext) treesitNodePrevSibling_autogen(node, named lispObject) (lispObject, error) {
	return ec.stub("treesit-node-prev-sibling") // Source file: treesit.c
}

func (ec *execContext) treesitNodeStart_autogen(node lispObject) (lispObject, error) {
	return ec.stub("treesit-node-start") // Source file: treesit.c
}

func (ec *execContext) treesitNodeString_autogen(node lispObject) (lispObject, error) {
	return ec.stub("treesit-node-string") // Source file: treesit.c
}

func (ec *execContext) treesitNodeType_autogen(node lispObject) (lispObject, error) {
	return ec.stub("treesit-node-type") // Source file: treesit.c
}

func (ec *execContext) treesitParserAddNotifier_autogen(parser, function lispObject) (lispObject, error) {
	return ec.stub("treesit-parser-add-notifier") // Source file: treesit.c
}

func (ec *execContext) treesitParserBuffer_autogen(parser lispObject) (lispObject, error) {
	return ec.stub("treesit-parser-buffer") // Source file: treesit.c
}

func (ec *execContext) treesitParserCreate_autogen(language, buffer, no_reuse lispObject) (lispObject, error) {
	return ec.stub("treesit-parser-create") // Source file: treesit.c
}

func (ec *execContext) treesitParserDelete_autogen(parser lispObject) (lispObject, error) {
	return ec.stub("treesit-parser-delete") // Source file: treesit.c
}

func (ec *execContext) treesitParserIncludedRanges_autogen(parser lispObject) (lispObject, error) {
	return ec.stub("treesit-parser-included-ranges") // Source file: treesit.c
}

func (ec *execContext) treesitParserLanguage_autogen(parser lispObject) (lispObject, error) {
	return ec.stub("treesit-parser-language") // Source file: treesit.c
}

func (ec *execContext) treesitParserList_autogen(buffer lispObject) (lispObject, error) {
	return ec.stub("treesit-parser-list") // Source file: treesit.c
}

func (ec *execContext) treesitParserNotifiers_autogen(parser lispObject) (lispObject, error) {
	return ec.stub("treesit-parser-notifiers") // Source file: treesit.c
}

func (ec *execContext) treesitParserP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("treesit-parser-p") // Source file: treesit.c
}

func (ec *execContext) treesitParserRemoveNotifier_autogen(parser, function lispObject) (lispObject, error) {
	return ec.stub("treesit-parser-remove-notifier") // Source file: treesit.c
}

func (ec *execContext) treesitParserRootNode_autogen(parser lispObject) (lispObject, error) {
	return ec.stub("treesit-parser-root-node") // Source file: treesit.c
}

func (ec *execContext) treesitParserSetIncludedRanges_autogen(parser, ranges lispObject) (lispObject, error) {
	return ec.stub("treesit-parser-set-included-ranges") // Source file: treesit.c
}

func (ec *execContext) treesitPatternExpand_autogen(pattern lispObject) (lispObject, error) {
	return ec.stub("treesit-pattern-expand") // Source file: treesit.c
}

func (ec *execContext) treesitQueryCapture_autogen(node, query, beg, end, node_only lispObject) (lispObject, error) {
	return ec.stub("treesit-query-capture") // Source file: treesit.c
}

func (ec *execContext) treesitQueryCompile_autogen(language, query, eager lispObject) (lispObject, error) {
	return ec.stub("treesit-query-compile") // Source file: treesit.c
}

func (ec *execContext) treesitQueryExpand_autogen(query lispObject) (lispObject, error) {
	return ec.stub("treesit-query-expand") // Source file: treesit.c
}

func (ec *execContext) treesitQueryLanguage_autogen(query lispObject) (lispObject, error) {
	return ec.stub("treesit-query-language") // Source file: treesit.c
}

func (ec *execContext) treesitQueryP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("treesit-query-p") // Source file: treesit.c
}

func (ec *execContext) treesitSearchForward_autogen(start, predicate, backward, all lispObject) (lispObject, error) {
	return ec.stub("treesit-search-forward") // Source file: treesit.c
}

func (ec *execContext) treesitSearchSubtree_autogen(node, predicate, backward, all, depth lispObject) (lispObject, error) {
	return ec.stub("treesit-search-subtree") // Source file: treesit.c
}

func (ec *execContext) treesitSubtreeStat_autogen(node lispObject) (lispObject, error) {
	return ec.stub("treesit-subtree-stat") // Source file: treesit.c
}

func (ec *execContext) undoBoundary_autogen() (lispObject, error) {
	return ec.stub("undo-boundary") // Source file: undo.c
}

func (ec *execContext) w16GetClipboardData_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("w16-get-clipboard-data") // Source file: w16select.c
}

func (ec *execContext) w16SelectionExistsP_autogen(selection, terminal lispObject) (lispObject, error) {
	return ec.stub("w16-selection-exists-p") // Source file: w16select.c
}

func (ec *execContext) w16SetClipboardData_autogen(string, frame lispObject) (lispObject, error) {
	return ec.stub("w16-set-clipboard-data") // Source file: w16select.c
}

func (ec *execContext) getScreenColor_autogen() (lispObject, error) {
	return ec.stub("get-screen-color") // Source file: w32console.c
}

func (ec *execContext) setCursorSize_autogen(size lispObject) (lispObject, error) {
	return ec.stub("set-cursor-size") // Source file: w32console.c
}

func (ec *execContext) setScreenColor_autogen(foreground, background lispObject) (lispObject, error) {
	return ec.stub("set-screen-color") // Source file: w32console.c
}

func (ec *execContext) w32BatteryStatus_autogen() (lispObject, error) {
	return ec.stub("w32-battery-status") // Source file: w32cygwinx.c
}

func (ec *execContext) defaultPrinterName_autogen() (lispObject, error) {
	return ec.stub("default-printer-name") // Source file: w32fns.c
}

func (ec *execContext) setMessageBeep_autogen(sound lispObject) (lispObject, error) {
	return ec.stub("set-message-beep") // Source file: w32fns.c
}

func (ec *execContext) systemMoveFileToTrash_autogen(filename lispObject) (lispObject, error) {
	return ec.stub("system-move-file-to-trash") // Source file: w32fns.c
}

func (ec *execContext) w32MenuBarInUse_autogen() (lispObject, error) {
	return ec.stub("w32--menu-bar-in-use") // Source file: w32fns.c
}

func (ec *execContext) w32DefineRgbColor_autogen(red, green, blue, name lispObject) (lispObject, error) {
	return ec.stub("w32-define-rgb-color") // Source file: w32fns.c
}

func (ec *execContext) w32DisplayMonitorAttributesList_autogen(display lispObject) (lispObject, error) {
	return ec.stub("w32-display-monitor-attributes-list") // Source file: w32fns.c
}

func (ec *execContext) w32FrameEdges_autogen(frame, type_ lispObject) (lispObject, error) {
	return ec.stub("w32-frame-edges") // Source file: w32fns.c
}

func (ec *execContext) w32FrameGeometry_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("w32-frame-geometry") // Source file: w32fns.c
}

func (ec *execContext) w32FrameListZOrder_autogen(display lispObject) (lispObject, error) {
	return ec.stub("w32-frame-list-z-order") // Source file: w32fns.c
}

func (ec *execContext) w32FrameRestack_autogen(frame1, frame2, above lispObject) (lispObject, error) {
	return ec.stub("w32-frame-restack") // Source file: w32fns.c
}

func (ec *execContext) w32GetImeOpenStatus_autogen() (lispObject, error) {
	return ec.stub("w32-get-ime-open-status") // Source file: w32fns.c
}

func (ec *execContext) w32MouseAbsolutePixelPosition_autogen() (lispObject, error) {
	return ec.stub("w32-mouse-absolute-pixel-position") // Source file: w32fns.c
}

func (ec *execContext) w32NotificationClose_autogen(id lispObject) (lispObject, error) {
	return ec.stub("w32-notification-close") // Source file: w32fns.c
}

func (ec *execContext) w32NotificationNotify_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("w32-notification-notify") // Source file: w32fns.c
}

func (ec *execContext) w32ReadRegistry_autogen(root, key, name lispObject) (lispObject, error) {
	return ec.stub("w32-read-registry") // Source file: w32fns.c
}

func (ec *execContext) w32ReconstructHotKey_autogen(hotkeyid lispObject) (lispObject, error) {
	return ec.stub("w32-reconstruct-hot-key") // Source file: w32fns.c
}

func (ec *execContext) w32RegisterHotKey_autogen(key lispObject) (lispObject, error) {
	return ec.stub("w32-register-hot-key") // Source file: w32fns.c
}

func (ec *execContext) w32RegisteredHotKeys_autogen() (lispObject, error) {
	return ec.stub("w32-registered-hot-keys") // Source file: w32fns.c
}

func (ec *execContext) w32SendSysCommand_autogen(command, frame lispObject) (lispObject, error) {
	return ec.stub("w32-send-sys-command") // Source file: w32fns.c
}

func (ec *execContext) w32SetImeOpenStatus_autogen(status lispObject) (lispObject, error) {
	return ec.stub("w32-set-ime-open-status") // Source file: w32fns.c
}

func (ec *execContext) w32SetMouseAbsolutePixelPosition_autogen(x, y lispObject) (lispObject, error) {
	return ec.stub("w32-set-mouse-absolute-pixel-position") // Source file: w32fns.c
}

func (ec *execContext) w32SetWallpaper_autogen(image_file lispObject) (lispObject, error) {
	return ec.stub("w32-set-wallpaper") // Source file: w32fns.c
}

func (ec *execContext) w32ShellExecute_autogen(operation, document, parameters, show_flag lispObject) (lispObject, error) {
	return ec.stub("w32-shell-execute") // Source file: w32fns.c
}

func (ec *execContext) w32ToggleLockKey_autogen(key, new_state lispObject) (lispObject, error) {
	return ec.stub("w32-toggle-lock-key") // Source file: w32fns.c
}

func (ec *execContext) w32UnregisterHotKey_autogen(key lispObject) (lispObject, error) {
	return ec.stub("w32-unregister-hot-key") // Source file: w32fns.c
}

func (ec *execContext) w32WindowExistsP_autogen(class, name lispObject) (lispObject, error) {
	return ec.stub("w32-window-exists-p") // Source file: w32fns.c
}

func (ec *execContext) xChangeWindowProperty_autogen(prop, value, frame, type_, format, outer_p lispObject) (lispObject, error) {
	return ec.stub("x-change-window-property") // Source file: w32fns.c
}

func (ec *execContext) xChangeWindowProperty_1_autogen(prop, value, frame, type_, format, outer_p, window_id lispObject) (lispObject, error) {
	return ec.stub("x-change-window-property") // Source file: xfns.c
}

func (ec *execContext) xDeleteWindowProperty_autogen(prop, frame lispObject) (lispObject, error) {
	return ec.stub("x-delete-window-property") // Source file: w32fns.c
}

func (ec *execContext) xDeleteWindowProperty_1_autogen(prop, frame, window_id lispObject) (lispObject, error) {
	return ec.stub("x-delete-window-property") // Source file: xfns.c
}

func (ec *execContext) xSynchronize_autogen(on, display lispObject) (lispObject, error) {
	return ec.stub("x-synchronize") // Source file: w32fns.c
}

func (ec *execContext) xSynchronize_1_autogen(on, terminal lispObject) (lispObject, error) {
	return ec.stub("x-synchronize") // Source file: xfns.c
}

func (ec *execContext) xWindowProperty_autogen(prop, frame, type_, source, delete_p, vector_ret_p lispObject) (lispObject, error) {
	return ec.stub("x-window-property") // Source file: w32fns.c
}

func (ec *execContext) xWindowProperty_1_autogen(prop, frame, type_, window_id, delete_p, vector_ret_p lispObject) (lispObject, error) {
	return ec.stub("x-window-property") // Source file: xfns.c
}

func (ec *execContext) w32notifyAddWatch_autogen(file, filter, callback lispObject) (lispObject, error) {
	return ec.stub("w32notify-add-watch") // Source file: w32notify.c
}

func (ec *execContext) w32notifyRmWatch_autogen(watch_descriptor lispObject) (lispObject, error) {
	return ec.stub("w32notify-rm-watch") // Source file: w32notify.c
}

func (ec *execContext) w32notifyValidP_autogen(watch_descriptor lispObject) (lispObject, error) {
	return ec.stub("w32notify-valid-p") // Source file: w32notify.c
}

func (ec *execContext) w32ApplicationType_autogen(program lispObject) (lispObject, error) {
	return ec.stub("w32-application-type") // Source file: w32proc.c
}

func (ec *execContext) w32GetCodepageCharset_autogen(cp lispObject) (lispObject, error) {
	return ec.stub("w32-get-codepage-charset") // Source file: w32proc.c
}

func (ec *execContext) w32GetConsoleCodepage_autogen() (lispObject, error) {
	return ec.stub("w32-get-console-codepage") // Source file: w32proc.c
}

func (ec *execContext) w32GetConsoleOutputCodepage_autogen() (lispObject, error) {
	return ec.stub("w32-get-console-output-codepage") // Source file: w32proc.c
}

func (ec *execContext) w32GetCurrentLocaleId_autogen() (lispObject, error) {
	return ec.stub("w32-get-current-locale-id") // Source file: w32proc.c
}

func (ec *execContext) w32GetDefaultLocaleId_autogen(userp lispObject) (lispObject, error) {
	return ec.stub("w32-get-default-locale-id") // Source file: w32proc.c
}

func (ec *execContext) w32GetKeyboardLayout_autogen() (lispObject, error) {
	return ec.stub("w32-get-keyboard-layout") // Source file: w32proc.c
}

func (ec *execContext) w32GetLocaleInfo_autogen(lcid, longform lispObject) (lispObject, error) {
	return ec.stub("w32-get-locale-info") // Source file: w32proc.c
}

func (ec *execContext) w32GetValidCodepages_autogen() (lispObject, error) {
	return ec.stub("w32-get-valid-codepages") // Source file: w32proc.c
}

func (ec *execContext) w32GetValidKeyboardLayouts_autogen() (lispObject, error) {
	return ec.stub("w32-get-valid-keyboard-layouts") // Source file: w32proc.c
}

func (ec *execContext) w32GetValidLocaleIds_autogen() (lispObject, error) {
	return ec.stub("w32-get-valid-locale-ids") // Source file: w32proc.c
}

func (ec *execContext) w32HasWinsock_autogen(load_now lispObject) (lispObject, error) {
	return ec.stub("w32-has-winsock") // Source file: w32proc.c
}

func (ec *execContext) w32LongFileName_autogen(filename lispObject) (lispObject, error) {
	return ec.stub("w32-long-file-name") // Source file: w32proc.c
}

func (ec *execContext) w32SetConsoleCodepage_autogen(cp lispObject) (lispObject, error) {
	return ec.stub("w32-set-console-codepage") // Source file: w32proc.c
}

func (ec *execContext) w32SetConsoleOutputCodepage_autogen(cp lispObject) (lispObject, error) {
	return ec.stub("w32-set-console-output-codepage") // Source file: w32proc.c
}

func (ec *execContext) w32SetCurrentLocale_autogen(lcid lispObject) (lispObject, error) {
	return ec.stub("w32-set-current-locale") // Source file: w32proc.c
}

func (ec *execContext) w32SetKeyboardLayout_autogen(layout lispObject) (lispObject, error) {
	return ec.stub("w32-set-keyboard-layout") // Source file: w32proc.c
}

func (ec *execContext) w32SetProcessPriority_autogen(process, priority lispObject) (lispObject, error) {
	return ec.stub("w32-set-process-priority") // Source file: w32proc.c
}

func (ec *execContext) w32ShortFileName_autogen(filename lispObject) (lispObject, error) {
	return ec.stub("w32-short-file-name") // Source file: w32proc.c
}

func (ec *execContext) w32UnloadWinsock_autogen() (lispObject, error) {
	return ec.stub("w32-unload-winsock") // Source file: w32proc.c
}

func (ec *execContext) w32GetClipboardData_autogen(ignored lispObject) (lispObject, error) {
	return ec.stub("w32-get-clipboard-data") // Source file: w32select.c
}

func (ec *execContext) w32SelectionExistsP_autogen(selection, terminal lispObject) (lispObject, error) {
	return ec.stub("w32-selection-exists-p") // Source file: w32select.c
}

func (ec *execContext) w32SelectionTargets_autogen(selection, terminal lispObject) (lispObject, error) {
	return ec.stub("w32-selection-targets") // Source file: w32select.c
}

func (ec *execContext) w32SetClipboardData_autogen(string, ignored lispObject) (lispObject, error) {
	return ec.stub("w32-set-clipboard-data") // Source file: w32select.c
}

func (ec *execContext) coordinatesInWindowP_autogen(coordinates, window lispObject) (lispObject, error) {
	return ec.stub("coordinates-in-window-p") // Source file: window.c
}

func (ec *execContext) currentWindowConfiguration_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("current-window-configuration") // Source file: window.c
}

func (ec *execContext) deleteOtherWindowsInternal_autogen(window, root lispObject) (lispObject, error) {
	return ec.stub("delete-other-windows-internal") // Source file: window.c
}

func (ec *execContext) deleteWindowInternal_autogen(window lispObject) (lispObject, error) {
	return ec.stub("delete-window-internal") // Source file: window.c
}

func (ec *execContext) forceWindowUpdate_autogen(object lispObject) (lispObject, error) {
	return ec.stub("force-window-update") // Source file: window.c
}

func (ec *execContext) frameFirstWindow_autogen(frame_or_window lispObject) (lispObject, error) {
	return ec.stub("frame-first-window") // Source file: window.c
}

func (ec *execContext) frameOldSelectedWindow_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("frame-old-selected-window") // Source file: window.c
}

func (ec *execContext) frameRootWindow_autogen(frame_or_window lispObject) (lispObject, error) {
	return ec.stub("frame-root-window") // Source file: window.c
}

func (ec *execContext) frameSelectedWindow_autogen(frame_or_window lispObject) (lispObject, error) {
	return ec.stub("frame-selected-window") // Source file: window.c
}

func (ec *execContext) getBufferWindow_autogen(buffer_or_name, all_frames lispObject) (lispObject, error) {
	return ec.stub("get-buffer-window") // Source file: window.c
}

func (ec *execContext) minibufferSelectedWindow_autogen() (lispObject, error) {
	return ec.stub("minibuffer-selected-window") // Source file: window.c
}

func (ec *execContext) minibufferWindow_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("minibuffer-window") // Source file: window.c
}

func (ec *execContext) moveToWindowLine_autogen(arg lispObject) (lispObject, error) {
	return ec.stub("move-to-window-line") // Source file: window.c
}

func (ec *execContext) nextWindow_autogen(window, minibuf, all_frames lispObject) (lispObject, error) {
	return ec.stub("next-window") // Source file: window.c
}

func (ec *execContext) oldSelectedWindow_autogen() (lispObject, error) {
	return ec.stub("old-selected-window") // Source file: window.c
}

func (ec *execContext) otherWindowForScrolling_autogen() (lispObject, error) {
	return ec.stub("other-window-for-scrolling") // Source file: window.c
}

func (ec *execContext) posVisibleInWindowP_autogen(pos, window, partially lispObject) (lispObject, error) {
	return ec.stub("pos-visible-in-window-p") // Source file: window.c
}

func (ec *execContext) previousWindow_autogen(window, minibuf, all_frames lispObject) (lispObject, error) {
	return ec.stub("previous-window") // Source file: window.c
}

func (ec *execContext) recenter_autogen(arg, redisplay lispObject) (lispObject, error) {
	return ec.stub("recenter") // Source file: window.c
}

func (ec *execContext) resizeMiniWindowInternal_autogen(window lispObject) (lispObject, error) {
	return ec.stub("resize-mini-window-internal") // Source file: window.c
}

func (ec *execContext) runWindowConfigurationChangeHook_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("run-window-configuration-change-hook") // Source file: window.c
}

func (ec *execContext) runWindowScrollFunctions_autogen(window lispObject) (lispObject, error) {
	return ec.stub("run-window-scroll-functions") // Source file: window.c
}

func (ec *execContext) scrollDown_autogen(arg lispObject) (lispObject, error) {
	return ec.stub("scroll-down") // Source file: window.c
}

func (ec *execContext) scrollLeft_autogen(arg, set_minimum lispObject) (lispObject, error) {
	return ec.stub("scroll-left") // Source file: window.c
}

func (ec *execContext) scrollRight_autogen(arg, set_minimum lispObject) (lispObject, error) {
	return ec.stub("scroll-right") // Source file: window.c
}

func (ec *execContext) scrollUp_autogen(arg lispObject) (lispObject, error) {
	return ec.stub("scroll-up") // Source file: window.c
}

func (ec *execContext) selectWindow_autogen(window, norecord lispObject) (lispObject, error) {
	return ec.stub("select-window") // Source file: window.c
}

func (ec *execContext) selectedWindow_autogen() (lispObject, error) {
	return ec.stub("selected-window") // Source file: window.c
}

func (ec *execContext) setFrameSelectedWindow_autogen(frame, window, norecord lispObject) (lispObject, error) {
	return ec.stub("set-frame-selected-window") // Source file: window.c
}

func (ec *execContext) setWindowBuffer_autogen(window, buffer_or_name, keep_margins lispObject) (lispObject, error) {
	return ec.stub("set-window-buffer") // Source file: window.c
}

func (ec *execContext) setWindowCombinationLimit_autogen(window, limit lispObject) (lispObject, error) {
	return ec.stub("set-window-combination-limit") // Source file: window.c
}

func (ec *execContext) setWindowConfiguration_autogen(configuration, dont_set_frame, dont_set_miniwindow lispObject) (lispObject, error) {
	return ec.stub("set-window-configuration") // Source file: window.c
}

func (ec *execContext) setWindowDedicatedP_autogen(window, flag lispObject) (lispObject, error) {
	return ec.stub("set-window-dedicated-p") // Source file: window.c
}

func (ec *execContext) setWindowDisplayTable_autogen(window, table lispObject) (lispObject, error) {
	return ec.stub("set-window-display-table") // Source file: window.c
}

func (ec *execContext) setWindowFringes_autogen(window, left_width, right_width, outside_margins, persistent lispObject) (lispObject, error) {
	return ec.stub("set-window-fringes") // Source file: window.c
}

func (ec *execContext) setWindowHscroll_autogen(window, ncol lispObject) (lispObject, error) {
	return ec.stub("set-window-hscroll") // Source file: window.c
}

func (ec *execContext) setWindowMargins_autogen(window, left_width, right_width lispObject) (lispObject, error) {
	return ec.stub("set-window-margins") // Source file: window.c
}

func (ec *execContext) setWindowNewNormal_autogen(window, size lispObject) (lispObject, error) {
	return ec.stub("set-window-new-normal") // Source file: window.c
}

func (ec *execContext) setWindowNewPixel_autogen(window, size, add lispObject) (lispObject, error) {
	return ec.stub("set-window-new-pixel") // Source file: window.c
}

func (ec *execContext) setWindowNewTotal_autogen(window, size, add lispObject) (lispObject, error) {
	return ec.stub("set-window-new-total") // Source file: window.c
}

func (ec *execContext) setWindowNextBuffers_autogen(window, next_buffers lispObject) (lispObject, error) {
	return ec.stub("set-window-next-buffers") // Source file: window.c
}

func (ec *execContext) setWindowParameter_autogen(window, parameter, value lispObject) (lispObject, error) {
	return ec.stub("set-window-parameter") // Source file: window.c
}

func (ec *execContext) setWindowPoint_autogen(window, pos lispObject) (lispObject, error) {
	return ec.stub("set-window-point") // Source file: window.c
}

func (ec *execContext) setWindowPrevBuffers_autogen(window, prev_buffers lispObject) (lispObject, error) {
	return ec.stub("set-window-prev-buffers") // Source file: window.c
}

func (ec *execContext) setWindowScrollBars_autogen(window, width, vertical_type, height, horizontal_type, persistent lispObject) (lispObject, error) {
	return ec.stub("set-window-scroll-bars") // Source file: window.c
}

func (ec *execContext) setWindowStart_autogen(window, pos, noforce lispObject) (lispObject, error) {
	return ec.stub("set-window-start") // Source file: window.c
}

func (ec *execContext) setWindowVscroll_autogen(window, vscroll, pixels_p, preserve_vscroll_p lispObject) (lispObject, error) {
	return ec.stub("set-window-vscroll") // Source file: window.c
}

func (ec *execContext) splitWindowInternal_autogen(old, pixel_size, side, normal_size lispObject) (lispObject, error) {
	return ec.stub("split-window-internal") // Source file: window.c
}

func (ec *execContext) windowAt_autogen(x, y, frame lispObject) (lispObject, error) {
	return ec.stub("window-at") // Source file: window.c
}

func (ec *execContext) windowBodyHeight_autogen(window, pixelwise lispObject) (lispObject, error) {
	return ec.stub("window-body-height") // Source file: window.c
}

func (ec *execContext) windowBodyWidth_autogen(window, pixelwise lispObject) (lispObject, error) {
	return ec.stub("window-body-width") // Source file: window.c
}

func (ec *execContext) windowBottomDividerWidth_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-bottom-divider-width") // Source file: window.c
}

func (ec *execContext) windowBuffer_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-buffer") // Source file: window.c
}

func (ec *execContext) windowBumpUseTime_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-bump-use-time") // Source file: window.c
}

func (ec *execContext) windowCombinationLimit_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-combination-limit") // Source file: window.c
}

func (ec *execContext) windowConfigurationEqualP_autogen(x, y lispObject) (lispObject, error) {
	return ec.stub("window-configuration-equal-p") // Source file: window.c
}

func (ec *execContext) windowConfigurationFrame_autogen(config lispObject) (lispObject, error) {
	return ec.stub("window-configuration-frame") // Source file: window.c
}

func (ec *execContext) windowConfigurationP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("window-configuration-p") // Source file: window.c
}

func (ec *execContext) windowDedicatedP_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-dedicated-p") // Source file: window.c
}

func (ec *execContext) windowDisplayTable_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-display-table") // Source file: window.c
}

func (ec *execContext) windowEnd_autogen(window, update lispObject) (lispObject, error) {
	return ec.stub("window-end") // Source file: window.c
}

func (ec *execContext) windowFrame_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-frame") // Source file: window.c
}

func (ec *execContext) windowFringes_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-fringes") // Source file: window.c
}

func (ec *execContext) windowHeaderLineHeight_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-header-line-height") // Source file: window.c
}

func (ec *execContext) windowHscroll_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-hscroll") // Source file: window.c
}

func (ec *execContext) windowLeftChild_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-left-child") // Source file: window.c
}

func (ec *execContext) windowLeftColumn_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-left-column") // Source file: window.c
}

func (ec *execContext) windowLineHeight_autogen(line, window lispObject) (lispObject, error) {
	return ec.stub("window-line-height") // Source file: window.c
}

func (ec *execContext) windowLinesPixelDimensions_autogen(window, first, last, body, inverse, left lispObject) (lispObject, error) {
	return ec.stub("window-lines-pixel-dimensions") // Source file: window.c
}

func (ec *execContext) windowList_autogen(frame, minibuf, window lispObject) (lispObject, error) {
	return ec.stub("window-list") // Source file: window.c
}

func (ec *execContext) windowList1_autogen(window, minibuf, all_frames lispObject) (lispObject, error) {
	return ec.stub("window-list-1") // Source file: window.c
}

func (ec *execContext) windowLiveP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("window-live-p") // Source file: window.c
}

func (ec *execContext) windowMargins_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-margins") // Source file: window.c
}

func (ec *execContext) windowMinibufferP_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-minibuffer-p") // Source file: window.c
}

func (ec *execContext) windowModeLineHeight_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-mode-line-height") // Source file: window.c
}

func (ec *execContext) windowNewNormal_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-new-normal") // Source file: window.c
}

func (ec *execContext) windowNewPixel_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-new-pixel") // Source file: window.c
}

func (ec *execContext) windowNewTotal_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-new-total") // Source file: window.c
}

func (ec *execContext) windowNextBuffers_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-next-buffers") // Source file: window.c
}

func (ec *execContext) windowNextSibling_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-next-sibling") // Source file: window.c
}

func (ec *execContext) windowNormalSize_autogen(window, horizontal lispObject) (lispObject, error) {
	return ec.stub("window-normal-size") // Source file: window.c
}

func (ec *execContext) windowOldBodyPixelHeight_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-old-body-pixel-height") // Source file: window.c
}

func (ec *execContext) windowOldBodyPixelWidth_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-old-body-pixel-width") // Source file: window.c
}

func (ec *execContext) windowOldBuffer_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-old-buffer") // Source file: window.c
}

func (ec *execContext) windowOldPixelHeight_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-old-pixel-height") // Source file: window.c
}

func (ec *execContext) windowOldPixelWidth_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-old-pixel-width") // Source file: window.c
}

func (ec *execContext) windowOldPoint_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-old-point") // Source file: window.c
}

func (ec *execContext) windowParameter_autogen(window, parameter lispObject) (lispObject, error) {
	return ec.stub("window-parameter") // Source file: window.c
}

func (ec *execContext) windowParameters_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-parameters") // Source file: window.c
}

func (ec *execContext) windowParent_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-parent") // Source file: window.c
}

func (ec *execContext) windowPixelHeight_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-pixel-height") // Source file: window.c
}

func (ec *execContext) windowPixelLeft_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-pixel-left") // Source file: window.c
}

func (ec *execContext) windowPixelTop_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-pixel-top") // Source file: window.c
}

func (ec *execContext) windowPixelWidth_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-pixel-width") // Source file: window.c
}

func (ec *execContext) windowPoint_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-point") // Source file: window.c
}

func (ec *execContext) windowPrevBuffers_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-prev-buffers") // Source file: window.c
}

func (ec *execContext) windowPrevSibling_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-prev-sibling") // Source file: window.c
}

func (ec *execContext) windowResizeApply_autogen(frame, horizontal lispObject) (lispObject, error) {
	return ec.stub("window-resize-apply") // Source file: window.c
}

func (ec *execContext) windowResizeApplyTotal_autogen(frame, horizontal lispObject) (lispObject, error) {
	return ec.stub("window-resize-apply-total") // Source file: window.c
}

func (ec *execContext) windowRightDividerWidth_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-right-divider-width") // Source file: window.c
}

func (ec *execContext) windowScrollBarHeight_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-scroll-bar-height") // Source file: window.c
}

func (ec *execContext) windowScrollBarWidth_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-scroll-bar-width") // Source file: window.c
}

func (ec *execContext) windowScrollBars_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-scroll-bars") // Source file: window.c
}

func (ec *execContext) windowStart_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-start") // Source file: window.c
}

func (ec *execContext) windowTabLineHeight_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-tab-line-height") // Source file: window.c
}

func (ec *execContext) windowTextHeight_autogen(window, pixelwise lispObject) (lispObject, error) {
	return ec.stub("window-text-height") // Source file: window.c
}

func (ec *execContext) windowTextWidth_autogen(window, pixelwise lispObject) (lispObject, error) {
	return ec.stub("window-text-width") // Source file: window.c
}

func (ec *execContext) windowTopChild_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-top-child") // Source file: window.c
}

func (ec *execContext) windowTopLine_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-top-line") // Source file: window.c
}

func (ec *execContext) windowTotalHeight_autogen(window, round lispObject) (lispObject, error) {
	return ec.stub("window-total-height") // Source file: window.c
}

func (ec *execContext) windowTotalWidth_autogen(window, round lispObject) (lispObject, error) {
	return ec.stub("window-total-width") // Source file: window.c
}

func (ec *execContext) windowUseTime_autogen(window lispObject) (lispObject, error) {
	return ec.stub("window-use-time") // Source file: window.c
}

func (ec *execContext) windowValidP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("window-valid-p") // Source file: window.c
}

func (ec *execContext) windowVscroll_autogen(window, pixels_p lispObject) (lispObject, error) {
	return ec.stub("window-vscroll") // Source file: window.c
}

func (ec *execContext) windowp_autogen(object lispObject) (lispObject, error) {
	return ec.stub("windowp") // Source file: window.c
}

func (ec *execContext) bidiFindOverriddenDirectionality_autogen(from, to, object, base_dir lispObject) (lispObject, error) {
	return ec.stub("bidi-find-overridden-directionality") // Source file: xdisp.c
}

func (ec *execContext) bidiResolvedLevels_autogen(vpos lispObject) (lispObject, error) {
	return ec.stub("bidi-resolved-levels") // Source file: xdisp.c
}

func (ec *execContext) bufferTextPixelSize_autogen(buffer_or_name, window, x_limit, y_limit lispObject) (lispObject, error) {
	return ec.stub("buffer-text-pixel-size") // Source file: xdisp.c
}

func (ec *execContext) currentBidiParagraphDirection_autogen(buffer lispObject) (lispObject, error) {
	return ec.stub("current-bidi-paragraph-direction") // Source file: xdisp.c
}

func (ec *execContext) displayLineIsContinuedP_autogen() (lispObject, error) {
	return ec.stub("display--line-is-continued-p") // Source file: xdisp.c
}

func (ec *execContext) dumpFrameGlyphMatrix_autogen() (lispObject, error) {
	return ec.stub("dump-frame-glyph-matrix") // Source file: xdisp.c
}

func (ec *execContext) dumpGlyphMatrix_autogen(glyphs lispObject) (lispObject, error) {
	return ec.stub("dump-glyph-matrix") // Source file: xdisp.c
}

func (ec *execContext) dumpGlyphRow_autogen(row, glyphs lispObject) (lispObject, error) {
	return ec.stub("dump-glyph-row") // Source file: xdisp.c
}

func (ec *execContext) dumpTabBarRow_autogen(row, glyphs lispObject) (lispObject, error) {
	return ec.stub("dump-tab-bar-row") // Source file: xdisp.c
}

func (ec *execContext) dumpToolBarRow_autogen(row, glyphs lispObject) (lispObject, error) {
	return ec.stub("dump-tool-bar-row") // Source file: xdisp.c
}

func (ec *execContext) formatModeLine_autogen(format, face, window, buffer lispObject) (lispObject, error) {
	return ec.stub("format-mode-line") // Source file: xdisp.c
}

func (ec *execContext) getDisplayProperty_autogen(position, prop, object, properties lispObject) (lispObject, error) {
	return ec.stub("get-display-property") // Source file: xdisp.c
}

func (ec *execContext) invisibleP_autogen(pos lispObject) (lispObject, error) {
	return ec.stub("invisible-p") // Source file: xdisp.c
}

func (ec *execContext) linePixelHeight_autogen() (lispObject, error) {
	return ec.stub("line-pixel-height") // Source file: xdisp.c
}

func (ec *execContext) longLineOptimizationsP_autogen() (lispObject, error) {
	return ec.stub("long-line-optimizations-p") // Source file: xdisp.c
}

func (ec *execContext) lookupImageMap_autogen(map_, x, y lispObject) (lispObject, error) {
	return ec.stub("lookup-image-map") // Source file: xdisp.c
}

func (ec *execContext) movePointVisually_autogen(direction lispObject) (lispObject, error) {
	return ec.stub("move-point-visually") // Source file: xdisp.c
}

func (ec *execContext) setBufferRedisplay_autogen(symbol, newval, op, where lispObject) (lispObject, error) {
	return ec.stub("set-buffer-redisplay") // Source file: xdisp.c
}

func (ec *execContext) tabBarHeight_autogen(frame, pixelwise lispObject) (lispObject, error) {
	return ec.stub("tab-bar-height") // Source file: xdisp.c
}

func (ec *execContext) toolBarHeight_autogen(frame, pixelwise lispObject) (lispObject, error) {
	return ec.stub("tool-bar-height") // Source file: xdisp.c
}

func (ec *execContext) traceRedisplay_autogen(arg lispObject) (lispObject, error) {
	return ec.stub("trace-redisplay") // Source file: xdisp.c
}

func (ec *execContext) traceToStderr_autogen(args ...lispObject) (lispObject, error) {
	return ec.stub("trace-to-stderr") // Source file: xdisp.c
}

func (ec *execContext) windowTextPixelSize_autogen(window, from, to, x_limit, y_limit, mode_lines, ignore_line_at_end lispObject) (lispObject, error) {
	return ec.stub("window-text-pixel-size") // Source file: xdisp.c
}

func (ec *execContext) bitmapSpecP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("bitmap-spec-p") // Source file: xfaces.c
}

func (ec *execContext) clearFaceCache_autogen(thoroughly lispObject) (lispObject, error) {
	return ec.stub("clear-face-cache") // Source file: xfaces.c
}

func (ec *execContext) colorDistance_autogen(color1, color2, frame, metric lispObject) (lispObject, error) {
	return ec.stub("color-distance") // Source file: xfaces.c
}

func (ec *execContext) colorGrayP_autogen(color, frame lispObject) (lispObject, error) {
	return ec.stub("color-gray-p") // Source file: xfaces.c
}

func (ec *execContext) colorSupportedP_autogen(color, frame, background_p lispObject) (lispObject, error) {
	return ec.stub("color-supported-p") // Source file: xfaces.c
}

func (ec *execContext) colorValuesFromColorSpec_autogen(spec lispObject) (lispObject, error) {
	return ec.stub("color-values-from-color-spec") // Source file: xfaces.c
}

func (ec *execContext) displaySupportsFaceAttributesP_autogen(attributes, display lispObject) (lispObject, error) {
	return ec.stub("display-supports-face-attributes-p") // Source file: xfaces.c
}

func (ec *execContext) dumpColors_autogen() (lispObject, error) {
	return ec.stub("dump-colors") // Source file: xfaces.c
}

func (ec *execContext) dumpFace_autogen(n lispObject) (lispObject, error) {
	return ec.stub("dump-face") // Source file: xfaces.c
}

func (ec *execContext) faceAttributeRelativeP_autogen(attribute, value lispObject) (lispObject, error) {
	return ec.stub("face-attribute-relative-p") // Source file: xfaces.c, attributes: const
}

func (ec *execContext) faceAttributesAsVector_autogen(plist lispObject) (lispObject, error) {
	return ec.stub("face-attributes-as-vector") // Source file: xfaces.c
}

func (ec *execContext) faceFont_autogen(face, frame, character lispObject) (lispObject, error) {
	return ec.stub("face-font") // Source file: xfaces.c
}

func (ec *execContext) frameFaceHashTable_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("frame--face-hash-table") // Source file: xfaces.c
}

func (ec *execContext) internalCopyLispFace_autogen(from, to, frame, new_frame lispObject) (lispObject, error) {
	return ec.stub("internal-copy-lisp-face") // Source file: xfaces.c
}

func (ec *execContext) internalFaceXGetResource_autogen(resource, class, frame lispObject) (lispObject, error) {
	return ec.stub("internal-face-x-get-resource") // Source file: xfaces.c
}

func (ec *execContext) internalGetLispFaceAttribute_autogen(symbol, keyword, frame lispObject) (lispObject, error) {
	return ec.stub("internal-get-lisp-face-attribute") // Source file: xfaces.c
}

func (ec *execContext) internalLispFaceAttributeValues_autogen(attr lispObject) (lispObject, error) {
	return ec.stub("internal-lisp-face-attribute-values") // Source file: xfaces.c
}

func (ec *execContext) internalLispFaceEmptyP_autogen(face, frame lispObject) (lispObject, error) {
	return ec.stub("internal-lisp-face-empty-p") // Source file: xfaces.c
}

func (ec *execContext) internalLispFaceEqualP_autogen(face1, face2, frame lispObject) (lispObject, error) {
	return ec.stub("internal-lisp-face-equal-p") // Source file: xfaces.c
}

func (ec *execContext) internalLispFaceP_autogen(face, frame lispObject) (lispObject, error) {
	return ec.stub("internal-lisp-face-p") // Source file: xfaces.c
}

func (ec *execContext) internalMakeLispFace_autogen(face, frame lispObject) (lispObject, error) {
	return ec.stub("internal-make-lisp-face") // Source file: xfaces.c
}

func (ec *execContext) internalMergeInGlobalFace_autogen(face, frame lispObject) (lispObject, error) {
	return ec.stub("internal-merge-in-global-face") // Source file: xfaces.c
}

func (ec *execContext) internalSetAlternativeFontFamilyAlist_autogen(alist lispObject) (lispObject, error) {
	return ec.stub("internal-set-alternative-font-family-alist") // Source file: xfaces.c
}

func (ec *execContext) internalSetAlternativeFontRegistryAlist_autogen(alist lispObject) (lispObject, error) {
	return ec.stub("internal-set-alternative-font-registry-alist") // Source file: xfaces.c
}

func (ec *execContext) internalSetFontSelectionOrder_autogen(order lispObject) (lispObject, error) {
	return ec.stub("internal-set-font-selection-order") // Source file: xfaces.c
}

func (ec *execContext) internalSetLispFaceAttribute_autogen(face, attr, value, frame lispObject) (lispObject, error) {
	return ec.stub("internal-set-lisp-face-attribute") // Source file: xfaces.c
}

func (ec *execContext) internalSetLispFaceAttributeFromResource_autogen(face, attr, value, frame lispObject) (lispObject, error) {
	return ec.stub("internal-set-lisp-face-attribute-from-resource") // Source file: xfaces.c
}

func (ec *execContext) mergeFaceAttribute_autogen(attribute, value1, value2 lispObject) (lispObject, error) {
	return ec.stub("merge-face-attribute") // Source file: xfaces.c
}

func (ec *execContext) showFaceResources_autogen() (lispObject, error) {
	return ec.stub("show-face-resources") // Source file: xfaces.c
}

func (ec *execContext) ttySuppressBoldInverseDefaultColors_autogen(suppress lispObject) (lispObject, error) {
	return ec.stub("tty-suppress-bold-inverse-default-colors") // Source file: xfaces.c
}

func (ec *execContext) xFamilyFonts_autogen(family, frame lispObject) (lispObject, error) {
	return ec.stub("x-family-fonts") // Source file: xfaces.c
}

func (ec *execContext) xListFonts_autogen(pattern, face, frame, maximum, width lispObject) (lispObject, error) {
	return ec.stub("x-list-fonts") // Source file: xfaces.c
}

func (ec *execContext) xLoadColorFile_autogen(filename lispObject) (lispObject, error) {
	return ec.stub("x-load-color-file") // Source file: xfaces.c
}

func (ec *execContext) xBackspaceDeleteKeysP_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("x-backspace-delete-keys-p") // Source file: xfns.c
}

func (ec *execContext) xBeginDrag_autogen(targets, action, frame, return_frame, allow_current_frame, follow_tooltip lispObject) (lispObject, error) {
	return ec.stub("x-begin-drag") // Source file: xfns.c
}

func (ec *execContext) xDisplayMonitorAttributesList_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-monitor-attributes-list") // Source file: xfns.c
}

func (ec *execContext) xDisplayLastUserTime_autogen(time_object, terminal lispObject) (lispObject, error) {
	return ec.stub("x-display-set-last-user-time") // Source file: xfns.c
}

func (ec *execContext) xFrameEdges_autogen(frame, type_ lispObject) (lispObject, error) {
	return ec.stub("x-frame-edges") // Source file: xfns.c
}

func (ec *execContext) xFrameGeometry_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("x-frame-geometry") // Source file: xfns.c
}

func (ec *execContext) xFrameListZOrder_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-frame-list-z-order") // Source file: xfns.c
}

func (ec *execContext) xFrameRestack_autogen(frame1, frame2, above lispObject) (lispObject, error) {
	return ec.stub("x-frame-restack") // Source file: xfns.c
}

func (ec *execContext) xGetModifierMasks_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-get-modifier-masks") // Source file: xfns.c
}

func (ec *execContext) xGetPageSetup_autogen() (lispObject, error) {
	return ec.stub("x-get-page-setup") // Source file: xfns.c
}

func (ec *execContext) xInternalFocusInputContext_autogen(focus lispObject) (lispObject, error) {
	return ec.stub("x-internal-focus-input-context") // Source file: xfns.c
}

func (ec *execContext) xMouseAbsolutePixelPosition_autogen() (lispObject, error) {
	return ec.stub("x-mouse-absolute-pixel-position") // Source file: xfns.c
}

func (ec *execContext) xPageSetupDialog_autogen() (lispObject, error) {
	return ec.stub("x-page-setup-dialog") // Source file: xfns.c
}

func (ec *execContext) xPrintFramesDialog_autogen(frames lispObject) (lispObject, error) {
	return ec.stub("x-print-frames-dialog") // Source file: xfns.c
}

func (ec *execContext) xServerInputExtensionVersion_autogen(terminal lispObject) (lispObject, error) {
	return ec.stub("x-server-input-extension-version") // Source file: xfns.c
}

func (ec *execContext) xSetMouseAbsolutePixelPosition_autogen(x, y lispObject) (lispObject, error) {
	return ec.stub("x-set-mouse-absolute-pixel-position") // Source file: xfns.c
}

func (ec *execContext) xTranslateCoordinates_autogen(frame, source_window, dest_window, source_x, source_y, require_child lispObject) (lispObject, error) {
	return ec.stub("x-translate-coordinates") // Source file: xfns.c
}

func (ec *execContext) xUsesOldGtkDialog_autogen() (lispObject, error) {
	return ec.stub("x-uses-old-gtk-dialog") // Source file: xfns.c
}

func (ec *execContext) xWindowPropertyAttributes_autogen(prop, frame, window_id lispObject) (lispObject, error) {
	return ec.stub("x-window-property-attributes") // Source file: xfns.c
}

func (ec *execContext) xWmSetSizeHint_autogen(frame lispObject) (lispObject, error) {
	return ec.stub("x-wm-set-size-hint") // Source file: xfns.c
}

func (ec *execContext) libxmlAvailableP_autogen() (lispObject, error) {
	return ec.stub("libxml-available-p") // Source file: xml.c
}

func (ec *execContext) libxmlParseHtmlRegion_autogen(start, end, base_url, discard_comments lispObject) (lispObject, error) {
	return ec.stub("libxml-parse-html-region") // Source file: xml.c
}

func (ec *execContext) libxmlParseXmlRegion_autogen(start, end, base_url, discard_comments lispObject) (lispObject, error) {
	return ec.stub("libxml-parse-xml-region") // Source file: xml.c
}

func (ec *execContext) xDisownSelectionInternal_autogen(selection, time_object, terminal lispObject) (lispObject, error) {
	return ec.stub("x-disown-selection-internal") // Source file: xselect.c
}

func (ec *execContext) xGetAtomName_autogen(value, frame lispObject) (lispObject, error) {
	return ec.stub("x-get-atom-name") // Source file: xselect.c
}

func (ec *execContext) xGetLocalSelection_autogen(value, target lispObject) (lispObject, error) {
	return ec.stub("x-get-local-selection") // Source file: xselect.c
}

func (ec *execContext) xGetSelectionInternal_autogen(selection_symbol, target_type, time_stamp, terminal lispObject) (lispObject, error) {
	return ec.stub("x-get-selection-internal") // Source file: xselect.c
}

func (ec *execContext) xOwnSelectionInternal_autogen(selection, value, frame lispObject) (lispObject, error) {
	return ec.stub("x-own-selection-internal") // Source file: xselect.c
}

func (ec *execContext) xRegisterDndAtom_autogen(atom, frame lispObject) (lispObject, error) {
	return ec.stub("x-register-dnd-atom") // Source file: xselect.c
}

func (ec *execContext) xSelectionExistsP_autogen(selection, terminal lispObject) (lispObject, error) {
	return ec.stub("x-selection-exists-p") // Source file: xselect.c
}

func (ec *execContext) xSelectionOwnerP_autogen(selection, terminal lispObject) (lispObject, error) {
	return ec.stub("x-selection-owner-p") // Source file: xselect.c
}

func (ec *execContext) xSendClientMessage_autogen(display, dest, from, message_type, format, values lispObject) (lispObject, error) {
	return ec.stub("x-send-client-message") // Source file: xselect.c
}

func (ec *execContext) toolBarGetSystemStyle_autogen() (lispObject, error) {
	return ec.stub("tool-bar-get-system-style") // Source file: xsettings.c
}

func (ec *execContext) handleSaveSession_autogen(event lispObject) (lispObject, error) {
	return ec.stub("handle-save-session") // Source file: xsmfns.c
}

func (ec *execContext) deleteXwidgetView_autogen(xwidget_view lispObject) (lispObject, error) {
	return ec.stub("delete-xwidget-view") // Source file: xwidget.c
}

func (ec *execContext) getBufferXwidgets_autogen(buffer lispObject) (lispObject, error) {
	return ec.stub("get-buffer-xwidgets") // Source file: xwidget.c
}

func (ec *execContext) killXwidget_autogen(xwidget lispObject) (lispObject, error) {
	return ec.stub("kill-xwidget") // Source file: xwidget.c
}

func (ec *execContext) makeXwidget_autogen(type_, title, width, height, arguments, buffer, related lispObject) (lispObject, error) {
	return ec.stub("make-xwidget") // Source file: xwidget.c
}

func (ec *execContext) setXwidgetBuffer_autogen(xwidget, buffer lispObject) (lispObject, error) {
	return ec.stub("set-xwidget-buffer") // Source file: xwidget.c
}

func (ec *execContext) setXwidgetPlist_autogen(xwidget, plist lispObject) (lispObject, error) {
	return ec.stub("set-xwidget-plist") // Source file: xwidget.c
}

func (ec *execContext) setXwidgetQueryOnExitFlag_autogen(xwidget, flag lispObject) (lispObject, error) {
	return ec.stub("set-xwidget-query-on-exit-flag") // Source file: xwidget.c
}

func (ec *execContext) xwidgetBuffer_autogen(xwidget lispObject) (lispObject, error) {
	return ec.stub("xwidget-buffer") // Source file: xwidget.c
}

func (ec *execContext) xwidgetInfo_autogen(xwidget lispObject) (lispObject, error) {
	return ec.stub("xwidget-info") // Source file: xwidget.c
}

func (ec *execContext) xwidgetLiveP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("xwidget-live-p") // Source file: xwidget.c
}

func (ec *execContext) xwidgetPerformLispyEvent_autogen(xwidget, event, frame lispObject) (lispObject, error) {
	return ec.stub("xwidget-perform-lispy-event") // Source file: xwidget.c
}

func (ec *execContext) xwidgetPlist_autogen(xwidget lispObject) (lispObject, error) {
	return ec.stub("xwidget-plist") // Source file: xwidget.c
}

func (ec *execContext) xwidgetQueryOnExitFlag_autogen(xwidget lispObject) (lispObject, error) {
	return ec.stub("xwidget-query-on-exit-flag") // Source file: xwidget.c
}

func (ec *execContext) xwidgetResize_autogen(xwidget, new_width, new_height lispObject) (lispObject, error) {
	return ec.stub("xwidget-resize") // Source file: xwidget.c
}

func (ec *execContext) xwidgetSizeRequest_autogen(xwidget lispObject) (lispObject, error) {
	return ec.stub("xwidget-size-request") // Source file: xwidget.c
}

func (ec *execContext) xwidgetViewInfo_autogen(xwidget_view lispObject) (lispObject, error) {
	return ec.stub("xwidget-view-info") // Source file: xwidget.c
}

func (ec *execContext) xwidgetViewLookup_autogen(xwidget, window lispObject) (lispObject, error) {
	return ec.stub("xwidget-view-lookup") // Source file: xwidget.c
}

func (ec *execContext) xwidgetViewModel_autogen(xwidget_view lispObject) (lispObject, error) {
	return ec.stub("xwidget-view-model") // Source file: xwidget.c
}

func (ec *execContext) xwidgetViewP_autogen(object lispObject) (lispObject, error) {
	return ec.stub("xwidget-view-p") // Source file: xwidget.c
}

func (ec *execContext) xwidgetViewWindow_autogen(xwidget_view lispObject) (lispObject, error) {
	return ec.stub("xwidget-view-window") // Source file: xwidget.c
}

func (ec *execContext) xwidgetWebkitBackForwardList_autogen(xwidget, limit lispObject) (lispObject, error) {
	return ec.stub("xwidget-webkit-back-forward-list") // Source file: xwidget.c
}

func (ec *execContext) xwidgetWebkitEstimatedLoadProgress_autogen(xwidget lispObject) (lispObject, error) {
	return ec.stub("xwidget-webkit-estimated-load-progress") // Source file: xwidget.c
}

func (ec *execContext) xwidgetWebkitExecuteScript_autogen(xwidget, script, fun lispObject) (lispObject, error) {
	return ec.stub("xwidget-webkit-execute-script") // Source file: xwidget.c
}

func (ec *execContext) xwidgetWebkitFinishSearch_autogen(xwidget lispObject) (lispObject, error) {
	return ec.stub("xwidget-webkit-finish-search") // Source file: xwidget.c
}

func (ec *execContext) xwidgetWebkitGotoHistory_autogen(xwidget, rel_pos lispObject) (lispObject, error) {
	return ec.stub("xwidget-webkit-goto-history") // Source file: xwidget.c
}

func (ec *execContext) xwidgetWebkitGotoUri_autogen(xwidget, uri lispObject) (lispObject, error) {
	return ec.stub("xwidget-webkit-goto-uri") // Source file: xwidget.c
}

func (ec *execContext) xwidgetWebkitLoadHtml_autogen(xwidget, text, base_uri lispObject) (lispObject, error) {
	return ec.stub("xwidget-webkit-load-html") // Source file: xwidget.c
}

func (ec *execContext) xwidgetWebkitNextResult_autogen(xwidget lispObject) (lispObject, error) {
	return ec.stub("xwidget-webkit-next-result") // Source file: xwidget.c
}

func (ec *execContext) xwidgetWebkitPreviousResult_autogen(xwidget lispObject) (lispObject, error) {
	return ec.stub("xwidget-webkit-previous-result") // Source file: xwidget.c
}

func (ec *execContext) xwidgetWebkitSearch_autogen(query, xwidget, case_insensitive, backwards, wrap_around lispObject) (lispObject, error) {
	return ec.stub("xwidget-webkit-search") // Source file: xwidget.c
}

func (ec *execContext) xwidgetWebkitSetCookieStorageFile_autogen(xwidget, file lispObject) (lispObject, error) {
	return ec.stub("xwidget-webkit-set-cookie-storage-file") // Source file: xwidget.c
}

func (ec *execContext) xwidgetWebkitStopLoading_autogen(xwidget lispObject) (lispObject, error) {
	return ec.stub("xwidget-webkit-stop-loading") // Source file: xwidget.c
}

func (ec *execContext) xwidgetWebkitTitle_autogen(xwidget lispObject) (lispObject, error) {
	return ec.stub("xwidget-webkit-title") // Source file: xwidget.c
}

func (ec *execContext) xwidgetWebkitUri_autogen(xwidget lispObject) (lispObject, error) {
	return ec.stub("xwidget-webkit-uri") // Source file: xwidget.c
}

func (ec *execContext) xwidgetWebkitZoom_autogen(xwidget, factor lispObject) (lispObject, error) {
	return ec.stub("xwidget-webkit-zoom") // Source file: xwidget.c
}

func (ec *execContext) xwidgetp_autogen(object lispObject) (lispObject, error) {
	return ec.stub("xwidgetp") // Source file: xwidget.c
}

func (ec *execContext) symbolsOfEmacs_autogen() {
	ec.defSubrM(nil, "bool-vector", ec.boolVector_autogen, 0)
	ec.defSubr2(nil, "cons", ec.cons_autogen, 2)
	ec.defSubr0(nil, "garbage-collect", ec.garbageCollect_autogen)
	ec.defSubr1(nil, "garbage-collect-maybe", ec.garbageCollectMaybe_autogen, 1)
	ec.defSubrM(nil, "list", ec.list_autogen, 0)
	ec.defSubr2(nil, "make-bool-vector", ec.makeBoolVector_autogen, 2)
	ec.defSubrM(nil, "make-byte-code", ec.makeByteCode_autogen, 4)
	ec.defSubrM(nil, "make-closure", ec.makeClosure_autogen, 1)
	ec.defSubr1(nil, "make-finalizer", ec.makeFinalizer_autogen, 1)
	ec.defSubr2(nil, "make-list", ec.makeList_autogen, 2)
	ec.defSubr0(nil, "make-marker", ec.makeMarker_autogen)
	ec.defSubr3(nil, "make-record", ec.makeRecord_autogen, 3)
	ec.defSubr3(nil, "make-string", ec.makeString_autogen, 2)
	ec.defSubr1(nil, "make-symbol", ec.makeSymbol_autogen, 1)
	ec.defSubr2(nil, "make-vector", ec.makeVector_autogen, 2)
	ec.defSubr0(nil, "malloc-info", ec.mallocInfo_autogen)
	ec.defSubr1(nil, "malloc-trim", ec.mallocTrim_autogen, 0)
	ec.defSubr0(nil, "memory-info", ec.memoryInfo_autogen)
	ec.defSubr0(nil, "memory-use-counts", ec.memoryUseCounts_autogen)
	ec.defSubr1(nil, "purecopy", ec.purecopy_autogen, 1)
	ec.defSubrM(nil, "record", ec.record_autogen, 1)
	ec.defSubr1(nil, "suspicious-object", ec.suspiciousObject_autogen, 1)
	ec.defSubrM(nil, "vector", ec.vector_autogen, 0)
	ec.defSubr0(nil, "debug-timer-check", ec.debugTimerCheck_autogen)
	ec.defSubr1(nil, "barf-if-buffer-read-only", ec.barfIfBufferReadOnly_autogen, 0)
	ec.defSubr1(nil, "buffer-base-buffer", ec.bufferBaseBuffer_autogen, 0)
	ec.defSubr1(nil, "buffer-chars-modified-tick", ec.bufferCharsModifiedTick_autogen, 0)
	ec.defSubr1(nil, "buffer-enable-undo", ec.bufferEnableUndo_autogen, 0)
	ec.defSubr1(nil, "buffer-file-name", ec.bufferFileName_autogen, 0)
	ec.defSubr1(nil, "buffer-list", ec.bufferList_autogen, 0)
	ec.defSubr1(nil, "buffer-live-p", ec.bufferLiveP_autogen, 1)
	ec.defSubr2(nil, "buffer-local-value", ec.bufferLocalValue_autogen, 2)
	ec.defSubr1(nil, "buffer-local-variables", ec.bufferLocalVariables_autogen, 0)
	ec.defSubr1(nil, "buffer-modified-p", ec.bufferModifiedP_autogen, 0)
	ec.defSubr1(nil, "buffer-modified-tick", ec.bufferModifiedTick_autogen, 0)
	ec.defSubr1(nil, "buffer-name", ec.bufferName_autogen, 0)
	ec.defSubr1(nil, "buffer-swap-text", ec.bufferSwapText_autogen, 1)
	ec.defSubr1(nil, "bury-buffer-internal", ec.buryBufferInternal_autogen, 1)
	ec.defSubr0(nil, "current-buffer", ec.currentBuffer_autogen)
	ec.defSubr1(nil, "delete-all-overlays", ec.deleteAllOverlays_autogen, 0)
	ec.defSubr1(nil, "delete-overlay", ec.deleteOverlay_autogen, 1)
	ec.defSubr0(nil, "erase-buffer", ec.eraseBuffer_autogen)
	ec.defSubr1(nil, "force-mode-line-update", ec.forceModeLineUpdate_autogen, 0)
	ec.defSubr2(nil, "generate-new-buffer-name", ec.generateNewBufferName_autogen, 1)
	ec.defSubr1(nil, "get-buffer", ec.getBuffer_autogen, 1)
	ec.defSubr2(nil, "get-buffer-create", ec.getBufferCreate_autogen, 1)
	ec.defSubr1(nil, "get-file-buffer", ec.getFileBuffer_autogen, 1)
	ec.defSubr2(nil, "internal--set-buffer-modified-tick", ec.internalSetBufferModifiedTick_autogen, 1)
	ec.defSubr1(nil, "kill-all-local-variables", ec.killAllLocalVariables_autogen, 0)
	ec.defSubr1(nil, "kill-buffer", ec.killBuffer_autogen, 0)
	ec.defSubr4(nil, "make-indirect-buffer", ec.makeIndirectBuffer_autogen, 2)
	ec.defSubr5(nil, "make-overlay", ec.makeOverlay_autogen, 2)
	ec.defSubr4(nil, "move-overlay", ec.moveOverlay_autogen, 3)
	ec.defSubr1(nil, "next-overlay-change", ec.nextOverlayChange_autogen, 1)
	ec.defSubr3(nil, "other-buffer", ec.otherBuffer_autogen, 0)
	ec.defSubr1(nil, "overlay-buffer", ec.overlayBuffer_autogen, 1)
	ec.defSubr1(nil, "overlay-end", ec.overlayEnd_autogen, 1)
	ec.defSubr2(nil, "overlay-get", ec.overlayGet_autogen, 2)
	ec.defSubr0(nil, "overlay-lists", ec.overlayLists_autogen)
	ec.defSubr1(nil, "overlay-properties", ec.overlayProperties_autogen, 1)
	ec.defSubr3(nil, "overlay-put", ec.overlayPut_autogen, 3)
	ec.defSubr1(nil, "overlay-recenter", ec.overlayRecenter_autogen, 1)
	ec.defSubr1(nil, "overlay-start", ec.overlayStart_autogen, 1)
	ec.defSubr1(nil, "overlay-tree", ec.overlayTree_autogen, 0)
	ec.defSubr1(nil, "overlayp", ec.overlayp_autogen, 1)
	ec.defSubr2(nil, "overlays-at", ec.overlaysAt_autogen, 1)
	ec.defSubr2(nil, "overlays-in", ec.overlaysIn_autogen, 2)
	ec.defSubr1(nil, "previous-overlay-change", ec.previousOverlayChange_autogen, 1)
	ec.defSubr2(nil, "rename-buffer", ec.renameBuffer_autogen, 1)
	ec.defSubr1(nil, "restore-buffer-modified-p", ec.restoreBufferModifiedP_autogen, 1)
	ec.defSubr1(nil, "set-buffer", ec.setBuffer_autogen, 1)
	ec.defSubr1(nil, "set-buffer-major-mode", ec.setBufferMajorMode_autogen, 1)
	ec.defSubr1(nil, "set-buffer-modified-p", ec.setBufferModifiedP_autogen, 1)
	ec.defSubr1(nil, "set-buffer-multibyte", ec.setBufferMultibyte_autogen, 1)
	ec.defSubr3(nil, "byte-code", ec.byteCode_autogen, 3)
	ec.defSubr0(nil, "internal-stack-stats", ec.internalStackStats_autogen)
	ec.defSubr3(nil, "call-interactively", ec.callInteractively_autogen, 1)
	ec.defSubrM(nil, "funcall-interactively", ec.funcallInteractively_autogen, 1)
	ec.defSubrU(nil, "interactive", ec.interactive_autogen, 0)
	ec.defSubr1(nil, "prefix-numeric-value", ec.prefixNumericValue_autogen, 1)
	ec.defSubrM(nil, "call-process", ec.callProcess_autogen, 1)
	ec.defSubrM(nil, "call-process-region", ec.callProcessRegion_autogen, 3)
	ec.defSubr2(nil, "getenv-internal", ec.getenvInternal_autogen, 1)
	ec.defSubr1(nil, "capitalize", ec.capitalize_autogen, 1)
	ec.defSubr3(nil, "capitalize-region", ec.capitalizeRegion_autogen, 2)
	ec.defSubr1(nil, "capitalize-word", ec.capitalizeWord_autogen, 1)
	ec.defSubr1(nil, "downcase", ec.downcase_autogen, 1)
	ec.defSubr3(nil, "downcase-region", ec.downcaseRegion_autogen, 2)
	ec.defSubr1(nil, "downcase-word", ec.downcaseWord_autogen, 1)
	ec.defSubr1(nil, "upcase", ec.upcase_autogen, 1)
	ec.defSubr1(nil, "upcase-initials", ec.upcaseInitials_autogen, 1)
	ec.defSubr3(nil, "upcase-initials-region", ec.upcaseInitialsRegion_autogen, 2)
	ec.defSubr3(nil, "upcase-region", ec.upcaseRegion_autogen, 2)
	ec.defSubr1(nil, "upcase-word", ec.upcaseWord_autogen, 1)
	ec.defSubr1(nil, "case-table-p", ec.caseTableP_autogen, 1)
	ec.defSubr0(nil, "current-case-table", ec.currentCaseTable_autogen)
	ec.defSubr1(nil, "set-case-table", ec.setCaseTable_autogen, 1)
	ec.defSubr1(nil, "set-standard-case-table", ec.setStandardCaseTable_autogen, 1)
	ec.defSubr0(nil, "standard-case-table", ec.standardCaseTable_autogen)
	ec.defSubr2(nil, "category-docstring", ec.categoryDocstring_autogen, 1)
	ec.defSubr1(nil, "category-set-mnemonics", ec.categorySetMnemonics_autogen, 1)
	ec.defSubr0(nil, "category-table", ec.categoryTable_autogen)
	ec.defSubr1(nil, "category-table-p", ec.categoryTableP_autogen, 1)
	ec.defSubr1(nil, "char-category-set", ec.charCategorySet_autogen, 1)
	ec.defSubr1(nil, "copy-category-table", ec.copyCategoryTable_autogen, 0)
	ec.defSubr3(nil, "define-category", ec.defineCategory_autogen, 2)
	ec.defSubr1(nil, "get-unused-category", ec.getUnusedCategory_autogen, 0)
	ec.defSubr1(nil, "make-category-set", ec.makeCategorySet_autogen, 1)
	ec.defSubr0(nil, "make-category-table", ec.makeCategoryTable_autogen)
	ec.defSubr4(nil, "modify-category-entry", ec.modifyCategoryEntry_autogen, 2)
	ec.defSubr1(nil, "set-category-table", ec.setCategoryTable_autogen, 1)
	ec.defSubr0(nil, "standard-category-table", ec.standardCategoryTable_autogen)
	ec.defSubr2(nil, "ccl-execute", ec.cclExecute_autogen, 2)
	ec.defSubr5(nil, "ccl-execute-on-string", ec.cclExecuteOnString_autogen, 3)
	ec.defSubr1(nil, "ccl-program-p", ec.cclProgramP_autogen, 1)
	ec.defSubr2(nil, "register-ccl-program", ec.registerCclProgram_autogen, 2)
	ec.defSubr2(nil, "register-code-conversion-map", ec.registerCodeConversionMap_autogen, 2)
	ec.defSubr1(nil, "char-resolve-modifiers", ec.charResolveModifiers_autogen, 1)
	ec.defSubr1(nil, "char-width", ec.charWidth_autogen, 1)
	ec.defSubr2(nil, "characterp", ec.characterp_autogen, 1)
	ec.defSubr2(nil, "get-byte", ec.getByte_autogen, 0)
	ec.defSubr1(nil, "max-char", ec.maxChar_autogen, 0)
	ec.defSubr1(nil, "multibyte-char-to-unibyte", ec.multibyteCharToUnibyte_autogen, 1)
	ec.defSubrM(nil, "string", ec.string_autogen, 0)
	ec.defSubr3(nil, "string-width", ec.stringWidth_autogen, 1)
	ec.defSubr1(nil, "unibyte-char-to-multibyte", ec.unibyteCharToMultibyte_autogen, 1)
	ec.defSubrM(nil, "unibyte-string", ec.unibyteString_autogen, 0)
	ec.defSubr2(nil, "char-charset", ec.charCharset_autogen, 1)
	ec.defSubr1(nil, "charset-after", ec.charsetAfter_autogen, 0)
	ec.defSubr1(nil, "charset-id-internal", ec.charsetIdInternal_autogen, 0)
	ec.defSubr1(nil, "charset-plist", ec.charsetPlist_autogen, 1)
	ec.defSubr1(nil, "charset-priority-list", ec.charsetPriorityList_autogen, 0)
	ec.defSubr1(nil, "charsetp", ec.charsetp_autogen, 1)
	ec.defSubr0(nil, "clear-charset-maps", ec.clearCharsetMaps_autogen)
	ec.defSubr4(nil, "declare-equiv-charset", ec.declareEquivCharset_autogen, 4)
	ec.defSubr2(nil, "decode-char", ec.decodeChar_autogen, 2)
	ec.defSubr2(nil, "define-charset-alias", ec.defineCharsetAlias_autogen, 2)
	ec.defSubrM(nil, "define-charset-internal", ec.defineCharsetInternal_autogen, emacsConstant_charset_arg_max_autogen)
	ec.defSubr2(nil, "encode-char", ec.encodeChar_autogen, 2)
	ec.defSubr3(nil, "find-charset-region", ec.findCharsetRegion_autogen, 2)
	ec.defSubr2(nil, "find-charset-string", ec.findCharsetString_autogen, 1)
	ec.defSubr2(nil, "get-unused-iso-final-char", ec.getUnusedIsoFinalChar_autogen, 2)
	ec.defSubr3(nil, "iso-charset", ec.isoCharset_autogen, 3)
	ec.defSubr5(nil, "make-char", ec.makeChar_autogen, 1)
	ec.defSubr5(nil, "map-charset-chars", ec.mapCharsetChars_autogen, 2)
	ec.defSubr2(nil, "set-charset-plist", ec.setCharsetPlist_autogen, 2)
	ec.defSubrM(nil, "set-charset-priority", ec.setCharsetPriority_autogen, 1)
	ec.defSubr1(nil, "sort-charsets", ec.sortCharsets_autogen, 1)
	ec.defSubr1(nil, "split-char", ec.splitChar_autogen, 1)
	ec.defSubr3(nil, "unify-charset", ec.unifyCharset_autogen, 1)
	ec.defSubr2(nil, "char-table-extra-slot", ec.charTableExtraSlot_autogen, 2)
	ec.defSubr1(nil, "char-table-parent", ec.charTableParent_autogen, 1)
	ec.defSubr2(nil, "char-table-range", ec.charTableRange_autogen, 2)
	ec.defSubr1(nil, "char-table-subtype", ec.charTableSubtype_autogen, 1)
	ec.defSubr2(nil, "get-unicode-property-internal", ec.getUnicodePropertyInternal_autogen, 2)
	ec.defSubr2(nil, "make-char-table", ec.makeCharTable_autogen, 1)
	ec.defSubr2(nil, "map-char-table", ec.mapCharTable_autogen, 2)
	ec.defSubr2(nil, "optimize-char-table", ec.optimizeCharTable_autogen, 1)
	ec.defSubr3(nil, "put-unicode-property-internal", ec.putUnicodePropertyInternal_autogen, 3)
	ec.defSubr3(nil, "set-char-table-extra-slot", ec.setCharTableExtraSlot_autogen, 3)
	ec.defSubr2(nil, "set-char-table-parent", ec.setCharTableParent_autogen, 2)
	ec.defSubr3(nil, "set-char-table-range", ec.setCharTableRange_autogen, 3)
	ec.defSubr1(nil, "unicode-property-table-internal", ec.unicodePropertyTableInternal_autogen, 1)
	ec.defSubr1(nil, "backward-char", ec.backwardChar_autogen, 0)
	ec.defSubr1(nil, "beginning-of-line", ec.beginningOfLine_autogen, 0)
	ec.defSubr2(nil, "delete-char", ec.deleteChar_autogen, 1)
	ec.defSubr1(nil, "end-of-line", ec.endOfLine_autogen, 0)
	ec.defSubr1(nil, "forward-char", ec.forwardChar_autogen, 0)
	ec.defSubr1(nil, "forward-line", ec.forwardLine_autogen, 0)
	ec.defSubr2(nil, "self-insert-command", ec.selfInsertCommand_autogen, 1)
	ec.defSubr1(nil, "check-coding-system", ec.checkCodingSystem_autogen, 1)
	ec.defSubr3(nil, "check-coding-systems-region", ec.checkCodingSystemsRegion_autogen, 3)
	ec.defSubr1(nil, "coding-system-aliases", ec.codingSystemAliases_autogen, 1)
	ec.defSubr1(nil, "coding-system-base", ec.codingSystemBase_autogen, 1)
	ec.defSubr1(nil, "coding-system-eol-type", ec.codingSystemEolType_autogen, 1)
	ec.defSubr1(nil, "coding-system-p", ec.codingSystemP_autogen, 1)
	ec.defSubr1(nil, "coding-system-plist", ec.codingSystemPlist_autogen, 1)
	ec.defSubr1(nil, "coding-system-priority-list", ec.codingSystemPriorityList_autogen, 0)
	ec.defSubr3(nil, "coding-system-put", ec.codingSystemPut_autogen, 3)
	ec.defSubr1(nil, "decode-big5-char", ec.decodeBig5Char_autogen, 1)
	ec.defSubr4(nil, "decode-coding-region", ec.decodeCodingRegion_autogen, 3)
	ec.defSubr4(nil, "decode-coding-string", ec.decodeCodingString_autogen, 2)
	ec.defSubr1(nil, "decode-sjis-char", ec.decodeSjisChar_autogen, 1)
	ec.defSubr2(nil, "define-coding-system-alias", ec.defineCodingSystemAlias_autogen, 2)
	ec.defSubrM(nil, "define-coding-system-internal", ec.defineCodingSystemInternal_autogen, emacsConstant_coding_arg_max_autogen)
	ec.defSubr3(nil, "detect-coding-region", ec.detectCodingRegion_autogen, 2)
	ec.defSubr2(nil, "detect-coding-string", ec.detectCodingString_autogen, 1)
	ec.defSubr1(nil, "encode-big5-char", ec.encodeBig5Char_autogen, 1)
	ec.defSubr4(nil, "encode-coding-region", ec.encodeCodingRegion_autogen, 3)
	ec.defSubr4(nil, "encode-coding-string", ec.encodeCodingString_autogen, 2)
	ec.defSubr1(nil, "encode-sjis-char", ec.encodeSjisChar_autogen, 1)
	ec.defSubr3(nil, "find-coding-systems-region-internal", ec.findCodingSystemsRegionInternal_autogen, 2)
	ec.defSubrM(nil, "find-operation-coding-system", ec.findOperationCodingSystem_autogen, 1)
	ec.defSubr7(nil, "internal-decode-string-utf-8", ec.internalDecodeStringUtf8_autogen, 7)
	ec.defSubr7(nil, "internal-encode-string-utf-8", ec.internalEncodeStringUtf8_autogen, 7)
	ec.defSubr1(nil, "keyboard-coding-system", ec.keyboardCodingSystem_autogen, 0)
	ec.defSubr2(nil, "read-coding-system", ec.readCodingSystem_autogen, 1)
	ec.defSubr1(nil, "read-non-nil-coding-system", ec.readNonNilCodingSystem_autogen, 1)
	ec.defSubrM(nil, "set-coding-system-priority", ec.setCodingSystemPriority_autogen, 0)
	ec.defSubr2(nil, "set-keyboard-coding-system-internal", ec.setKeyboardCodingSystemInternal_autogen, 1)
	ec.defSubr1(nil, "set-safe-terminal-coding-system-internal", ec.setSafeTerminalCodingSystemInternal_autogen, 1)
	ec.defSubr2(nil, "set-terminal-coding-system-internal", ec.setTerminalCodingSystemInternal_autogen, 1)
	ec.defSubr1(nil, "terminal-coding-system", ec.terminalCodingSystem_autogen, 0)
	ec.defSubr5(nil, "unencodable-char-position", ec.unencodableCharPosition_autogen, 3)
	ec.defSubr1(nil, "comp--compile-ctxt-to-file", ec.compCompileCtxtToFile_autogen, 1)
	ec.defSubr0(nil, "comp--init-ctxt", ec.compInitCtxt_autogen)
	ec.defSubr2(nil, "comp--install-trampoline", ec.compInstallTrampoline_autogen, 2)
	ec.defSubr7(nil, "comp--late-register-subr", ec.compLateRegisterSubr_autogen, 7)
	ec.defSubr7(nil, "comp--register-lambda", ec.compRegisterLambda_autogen, 7)
	ec.defSubr7(nil, "comp--register-subr", ec.compRegisterSubr_autogen, 7)
	ec.defSubr0(nil, "comp--release-ctxt", ec.compReleaseCtxt_autogen)
	ec.defSubr1(nil, "comp--subr-signature", ec.compSubrSignature_autogen, 1)
	ec.defSubr2(nil, "comp-el-to-eln-filename", ec.compElToElnFilename_autogen, 1)
	ec.defSubr1(nil, "comp-el-to-eln-rel-filename", ec.compElToElnRelFilename_autogen, 1)
	ec.defSubr0(nil, "comp-libgccjit-version", ec.compLibgccjitVersion_autogen)
	ec.defSubr0(nil, "comp-native-compiler-options-effective-p", ec.compNativeCompilerOptionsEffectiveP_autogen)
	ec.defSubr0(nil, "comp-native-driver-options-effective-p", ec.compNativeDriverOptionsEffectiveP_autogen)
	ec.defSubr0(nil, "native-comp-available-p", ec.nativeCompAvailableP_autogen)
	ec.defSubr2(nil, "native-elisp-load", ec.nativeElispLoad_autogen, 1)
	ec.defSubr0(nil, "clear-composition-cache", ec.clearCompositionCache_autogen)
	ec.defSubr4(nil, "compose-region-internal", ec.composeRegionInternal_autogen, 2)
	ec.defSubr5(nil, "compose-string-internal", ec.composeStringInternal_autogen, 3)
	ec.defSubr4(nil, "composition-get-gstring", ec.compositionGetGstring_autogen, 4)
	ec.defSubr1(nil, "composition-sort-rules", ec.compositionSortRules_autogen, 1)
	ec.defSubr4(nil, "find-composition-internal", ec.findCompositionInternal_autogen, 4)
	ec.defSubr2(nil, "cygwin-convert-file-name-from-windows", ec.cygwinConvertFileNameFromWindows_autogen, 1)
	ec.defSubr2(nil, "cygwin-convert-file-name-to-windows", ec.cygwinConvertFileNameToWindows_autogen, 1)
	ec.defSubr2(nil, "%", ec.rem_autogen, 2)
	ec.defSubrM(nil, "*", ec.times_autogen, 0)
	ec.defSubrM(nil, "+", ec.plus_autogen, 0)
	ec.defSubrM(nil, "-", ec.minus_autogen, 0)
	ec.defSubrM(nil, "/", ec.quo_autogen, 1)
	ec.defSubr2(nil, "/=", ec.neq_autogen, 2)
	ec.defSubr1(nil, "1+", ec.add1_autogen, 1)
	ec.defSubr1(nil, "1-", ec.sub1_autogen, 1)
	ec.defSubrM(nil, "<", ec.lss_autogen, 1)
	ec.defSubrM(nil, "<=", ec.leq_autogen, 1)
	ec.defSubrM(nil, "=", ec.eqlsign_autogen, 1)
	ec.defSubrM(nil, ">", ec.gtr_autogen, 1)
	ec.defSubrM(nil, ">=", ec.geq_autogen, 1)
	ec.defSubr2(nil, "add-variable-watcher", ec.addVariableWatcher_autogen, 2)
	ec.defSubr2(nil, "aref", ec.aref_autogen, 2)
	ec.defSubr1(nil, "arrayp", ec.arrayp_autogen, 1)
	ec.defSubr3(nil, "aset", ec.aset_autogen, 3)
	ec.defSubr2(nil, "ash", ec.ash_autogen, 2)
	ec.defSubr1(nil, "atom", ec.atom_autogen, 1)
	ec.defSubr1(nil, "bare-symbol", ec.bareSymbol_autogen, 1)
	ec.defSubr1(nil, "bare-symbol-p", ec.bareSymbolP_autogen, 1)
	ec.defSubr3(nil, "bool-vector-count-consecutive", ec.boolVectorCountConsecutive_autogen, 3)
	ec.defSubr1(nil, "bool-vector-count-population", ec.boolVectorCountPopulation_autogen, 1)
	ec.defSubr3(nil, "bool-vector-exclusive-or", ec.boolVectorExclusiveOr_autogen, 2)
	ec.defSubr3(nil, "bool-vector-intersection", ec.boolVectorIntersection_autogen, 2)
	ec.defSubr2(nil, "bool-vector-not", ec.boolVectorNot_autogen, 1)
	ec.defSubr1(nil, "bool-vector-p", ec.boolVectorP_autogen, 1)
	ec.defSubr3(nil, "bool-vector-set-difference", ec.boolVectorSetDifference_autogen, 2)
	ec.defSubr2(nil, "bool-vector-subsetp", ec.boolVectorSubsetp_autogen, 2)
	ec.defSubr3(nil, "bool-vector-union", ec.boolVectorUnion_autogen, 2)
	ec.defSubr1(nil, "boundp", ec.boundp_autogen, 1)
	ec.defSubr1(nil, "bufferp", ec.bufferp_autogen, 1)
	ec.defSubr1(nil, "byte-code-function-p", ec.byteCodeFunctionP_autogen, 1)
	ec.defSubr0(nil, "byteorder", ec.byteorder_autogen)
	ec.defSubr1(nil, "car", ec.car_autogen, 1)
	ec.defSubr1(nil, "car-safe", ec.carSafe_autogen, 1)
	ec.defSubr1(nil, "cdr", ec.cdr_autogen, 1)
	ec.defSubr1(nil, "cdr-safe", ec.cdrSafe_autogen, 1)
	ec.defSubr1(nil, "char-or-string-p", ec.charOrStringP_autogen, 1)
	ec.defSubr1(nil, "char-table-p", ec.charTableP_autogen, 1)
	ec.defSubr1(nil, "command-modes", ec.commandModes_autogen, 1)
	ec.defSubr1(nil, "condition-variable-p", ec.conditionVariableP_autogen, 1)
	ec.defSubr1(nil, "consp", ec.consp_autogen, 1)
	ec.defSubr3(nil, "defalias", ec.defalias_autogen, 2)
	ec.defSubr1(nil, "default-boundp", ec.defaultBoundp_autogen, 1)
	ec.defSubr1(nil, "default-value", ec.defaultValue_autogen, 1)
	ec.defSubr2(nil, "eq", ec.eq_autogen, 2)
	ec.defSubr1(nil, "fboundp", ec.fboundp_autogen, 1)
	ec.defSubr1(nil, "floatp", ec.floatp_autogen, 1)
	ec.defSubr1(nil, "fmakunbound", ec.fmakunbound_autogen, 1)
	ec.defSubr2(nil, "fset", ec.fset_autogen, 2)
	ec.defSubr1(nil, "get-variable-watchers", ec.getVariableWatchers_autogen, 1)
	ec.defSubr2(nil, "indirect-function", ec.indirectFunction_autogen, 1)
	ec.defSubr1(nil, "indirect-variable", ec.indirectVariable_autogen, 1)
	ec.defSubr1(nil, "integer-or-marker-p", ec.integerOrMarkerP_autogen, 1)
	ec.defSubr1(nil, "integerp", ec.integerp_autogen, 1)
	ec.defSubr1(nil, "interactive-form", ec.interactiveForm_autogen, 1)
	ec.defSubr1(nil, "keywordp", ec.keywordp_autogen, 1)
	ec.defSubr1(nil, "kill-local-variable", ec.killLocalVariable_autogen, 1)
	ec.defSubr1(nil, "listp", ec.listp_autogen, 1)
	ec.defSubr2(nil, "local-variable-if-set-p", ec.localVariableIfSetP_autogen, 1)
	ec.defSubr2(nil, "local-variable-p", ec.localVariableP_autogen, 1)
	ec.defSubrM(nil, "logand", ec.logand_autogen, 0)
	ec.defSubr1(nil, "logcount", ec.logcount_autogen, 1)
	ec.defSubrM(nil, "logior", ec.logior_autogen, 0)
	ec.defSubr1(nil, "lognot", ec.lognot_autogen, 1)
	ec.defSubrM(nil, "logxor", ec.logxor_autogen, 0)
	ec.defSubr1(nil, "make-local-variable", ec.makeLocalVariable_autogen, 1)
	ec.defSubr1(nil, "make-variable-buffer-local", ec.makeVariableBufferLocal_autogen, 1)
	ec.defSubr1(nil, "makunbound", ec.makunbound_autogen, 1)
	ec.defSubr1(nil, "markerp", ec.markerp_autogen, 1)
	ec.defSubrM(nil, "max", ec.max_autogen, 1)
	ec.defSubrM(nil, "min", ec.min_autogen, 1)
	ec.defSubr2(nil, "mod", ec.mod_autogen, 2)
	ec.defSubr1(nil, "module-function-p", ec.moduleFunctionP_autogen, 1)
	ec.defSubr1(nil, "multibyte-string-p", ec.multibyteStringP_autogen, 1)
	ec.defSubr1(nil, "mutexp", ec.mutexp_autogen, 1)
	ec.defSubr1(nil, "native-comp-unit-file", ec.nativeCompUnitFile_autogen, 1)
	ec.defSubr2(nil, "native-comp-unit-set-file", ec.nativeCompUnitSetFile_autogen, 2)
	ec.defSubr1(nil, "natnump", ec.natnump_autogen, 1)
	ec.defSubr1(nil, "nlistp", ec.nlistp_autogen, 1)
	ec.defSubr1(nil, "null", ec.null_autogen, 1)
	ec.defSubr1(nil, "number-or-marker-p", ec.numberOrMarkerP_autogen, 1)
	ec.defSubr1(nil, "number-to-string", ec.numberToString_autogen, 1)
	ec.defSubr1(nil, "numberp", ec.numberp_autogen, 1)
	ec.defSubr2(nil, "position-symbol", ec.positionSymbol_autogen, 2)
	ec.defSubr1(nil, "recordp", ec.recordp_autogen, 1)
	ec.defSubr1(nil, "remove-pos-from-symbol", ec.removePosFromSymbol_autogen, 1)
	ec.defSubr2(nil, "remove-variable-watcher", ec.removeVariableWatcher_autogen, 2)
	ec.defSubr1(nil, "sequencep", ec.sequencep_autogen, 1)
	ec.defSubr2(nil, "set", ec.set_autogen, 2)
	ec.defSubr2(nil, "set-default", ec.setDefault_autogen, 2)
	ec.defSubr2(nil, "setcar", ec.setcar_autogen, 2)
	ec.defSubr2(nil, "setcdr", ec.setcdr_autogen, 2)
	ec.defSubr2(nil, "setplist", ec.setplist_autogen, 2)
	ec.defSubr2(nil, "string-to-number", ec.stringToNumber_autogen, 1)
	ec.defSubr1(nil, "stringp", ec.stringp_autogen, 1)
	ec.defSubr1(nil, "subr-arity", ec.subrArity_autogen, 1)
	ec.defSubr1(nil, "subr-name", ec.subrName_autogen, 1)
	ec.defSubr1(nil, "subr-native-comp-unit", ec.subrNativeCompUnit_autogen, 1)
	ec.defSubr1(nil, "subr-native-elisp-p", ec.subrNativeElispP_autogen, 1)
	ec.defSubr1(nil, "subr-native-lambda-list", ec.subrNativeLambdaList_autogen, 1)
	ec.defSubr1(nil, "subr-type", ec.subrType_autogen, 1)
	ec.defSubr1(nil, "subrp", ec.subrp_autogen, 1)
	ec.defSubr1(nil, "symbol-function", ec.symbolFunction_autogen, 1)
	ec.defSubr1(nil, "symbol-name", ec.symbolName_autogen, 1)
	ec.defSubr1(nil, "symbol-plist", ec.symbolPlist_autogen, 1)
	ec.defSubr1(nil, "symbol-value", ec.symbolValue_autogen, 1)
	ec.defSubr1(nil, "symbol-with-pos-p", ec.symbolWithPosP_autogen, 1)
	ec.defSubr1(nil, "symbol-with-pos-pos", ec.symbolWithPosPos_autogen, 1)
	ec.defSubr1(nil, "symbolp", ec.symbolp_autogen, 1)
	ec.defSubr1(nil, "threadp", ec.threadp_autogen, 1)
	ec.defSubr1(nil, "type-of", ec.typeOf_autogen, 1)
	ec.defSubr1(nil, "user-ptrp", ec.userPtrp_autogen, 1)
	ec.defSubr1(nil, "variable-binding-locus", ec.variableBindingLocus_autogen, 1)
	ec.defSubr1(nil, "vector-or-char-table-p", ec.vectorOrCharTableP_autogen, 1)
	ec.defSubr1(nil, "vectorp", ec.vectorp_autogen, 1)
	ec.defSubr2(nil, "dbus--init-bus", ec.dbusInitBus_autogen, 1)
	ec.defSubr1(nil, "dbus-get-unique-name", ec.dbusGetUniqueName_autogen, 1)
	ec.defSubrM(nil, "dbus-message-internal", ec.dbusMessageInternal_autogen, 4)
	ec.defSubr0(nil, "zlib-available-p", ec.zlibAvailableP_autogen)
	ec.defSubr3(nil, "zlib-decompress-region", ec.zlibDecompressRegion_autogen, 2)
	ec.defSubr5(nil, "directory-files", ec.directoryFiles_autogen, 1)
	ec.defSubr6(nil, "directory-files-and-attributes", ec.directoryFilesAndAttributes_autogen, 1)
	ec.defSubr2(nil, "file-attributes", ec.fileAttributes_autogen, 1)
	ec.defSubr2(nil, "file-attributes-lessp", ec.fileAttributesLessp_autogen, 2)
	ec.defSubr2(nil, "file-name-all-completions", ec.fileNameAllCompletions_autogen, 2)
	ec.defSubr3(nil, "file-name-completion", ec.fileNameCompletion_autogen, 2)
	ec.defSubr0(nil, "system-groups", ec.systemGroups_autogen)
	ec.defSubr0(nil, "system-users", ec.systemUsers_autogen)
	ec.defSubr1(nil, "ding", ec.ding_autogen, 0)
	ec.defSubr2(nil, "display--update-for-mouse-movement", ec.displayUpdateForMouseMovement_autogen, 2)
	ec.defSubr0(nil, "dump-redisplay-history", ec.dumpRedisplayHistory_autogen)
	ec.defSubr1(nil, "frame-or-buffer-changed-p", ec.frameOrBufferChangedP_autogen, 0)
	ec.defSubr2(nil, "internal-show-cursor", ec.internalShowCursor_autogen, 2)
	ec.defSubr1(nil, "internal-show-cursor-p", ec.internalShowCursorP_autogen, 0)
	ec.defSubr1(nil, "open-termscript", ec.openTermscript_autogen, 1)
	ec.defSubr1(nil, "redisplay", ec.redisplay_autogen, 0)
	ec.defSubr0(nil, "redraw-display", ec.redrawDisplay_autogen)
	ec.defSubr1(nil, "redraw-frame", ec.redrawFrame_autogen, 0)
	ec.defSubr2(nil, "send-string-to-terminal", ec.sendStringToTerminal_autogen, 1)
	ec.defSubr2(nil, "sleep-for", ec.sleepFor_autogen, 1)
	ec.defSubr1(nil, "Snarf-documentation", ec.snarfDocumentation_autogen, 1)
	ec.defSubr2(nil, "documentation", ec.documentation_autogen, 1)
	ec.defSubr3(nil, "documentation-property", ec.documentationProperty_autogen, 2)
	ec.defSubr0(nil, "text-quoting-style", ec.textQuotingStyle_autogen)
	ec.defSubr1(nil, "file-system-info", ec.fileSystemInfo_autogen, 1)
	ec.defSubr1(nil, "file-system-info", ec.fileSystemInfo_1_autogen, 1)
	ec.defSubr1(nil, "file-system-info", ec.fileSystemInfo_2_autogen, 1)
	ec.defSubr0(nil, "insert-startup-screen", ec.insertStartupScreen_autogen)
	ec.defSubr2(nil, "int86", ec.int86_autogen, 2)
	ec.defSubr2(nil, "msdos-memget", ec.dosMemget_autogen, 2)
	ec.defSubr2(nil, "msdos-memput", ec.dosMemput_autogen, 2)
	ec.defSubr0(nil, "msdos-mouse-disable", ec.msdosMouseDisable_autogen)
	ec.defSubr0(nil, "msdos-mouse-enable", ec.msdosMouseEnable_autogen)
	ec.defSubr0(nil, "msdos-mouse-init", ec.msdosMouseInit_autogen)
	ec.defSubr0(nil, "msdos-mouse-p", ec.msdosMouseP_autogen)
	ec.defSubr2(nil, "msdos-set-keyboard", ec.msdosSetKeyboard_autogen, 1)
	ec.defSubr0(nil, "bobp", ec.bobp_autogen)
	ec.defSubr0(nil, "bolp", ec.bolp_autogen)
	ec.defSubr1(nil, "buffer-size", ec.bufferSize_autogen, 0)
	ec.defSubr0(nil, "buffer-string", ec.bufferString_autogen)
	ec.defSubr2(nil, "buffer-substring", ec.bufferSubstring_autogen, 2)
	ec.defSubr2(nil, "buffer-substring-no-properties", ec.bufferSubstringNoProperties_autogen, 2)
	ec.defSubr1(nil, "byte-to-position", ec.byteToPosition_autogen, 1)
	ec.defSubr1(nil, "byte-to-string", ec.byteToString_autogen, 1)
	ec.defSubr1(nil, "char-after", ec.charAfter_autogen, 0)
	ec.defSubr1(nil, "char-before", ec.charBefore_autogen, 0)
	ec.defSubr2(nil, "char-equal", ec.charEqual_autogen, 2)
	ec.defSubr1(nil, "char-to-string", ec.charToString_autogen, 1)
	ec.defSubr6(nil, "compare-buffer-substrings", ec.compareBufferSubstrings_autogen, 6)
	ec.defSubr5(nil, "constrain-to-field", ec.constrainToField_autogen, 2)
	ec.defSubr0(nil, "current-message", ec.currentMessage_autogen)
	ec.defSubr2(nil, "delete-and-extract-region", ec.deleteAndExtractRegion_autogen, 2)
	ec.defSubr1(nil, "delete-field", ec.deleteField_autogen, 0)
	ec.defSubr2(nil, "delete-region", ec.deleteRegion_autogen, 2)
	ec.defSubr0(nil, "emacs-pid", ec.emacsPid_autogen)
	ec.defSubr0(nil, "eobp", ec.eobp_autogen)
	ec.defSubr0(nil, "eolp", ec.eolp_autogen)
	ec.defSubr3(nil, "field-beginning", ec.fieldBeginning_autogen, 0)
	ec.defSubr3(nil, "field-end", ec.fieldEnd_autogen, 0)
	ec.defSubr1(nil, "field-string", ec.fieldString_autogen, 0)
	ec.defSubr1(nil, "field-string-no-properties", ec.fieldStringNoProperties_autogen, 0)
	ec.defSubr0(nil, "following-char", ec.followingChar_autogen)
	ec.defSubrM(nil, "format", ec.format_autogen, 1)
	ec.defSubrM(nil, "format-message", ec.formatMessage_autogen, 1)
	ec.defSubr0(nil, "gap-position", ec.gapPosition_autogen)
	ec.defSubr0(nil, "gap-size", ec.gapSize_autogen)
	ec.defSubr3(nil, "get-pos-property", ec.getPosProperty_autogen, 2)
	ec.defSubr1(nil, "goto-char", ec.gotoChar_autogen, 1)
	ec.defSubr0(nil, "group-gid", ec.groupGid_autogen)
	ec.defSubr1(nil, "group-name", ec.groupName_autogen, 1)
	ec.defSubr0(nil, "group-real-gid", ec.groupRealGid_autogen)
	ec.defSubrM(nil, "insert", ec.insert_autogen, 0)
	ec.defSubrM(nil, "insert-and-inherit", ec.insertAndInherit_autogen, 0)
	ec.defSubrM(nil, "insert-before-markers", ec.insertBeforeMarkers_autogen, 0)
	ec.defSubrM(nil, "insert-before-markers-and-inherit", ec.insertAndInheritBeforeMarkers_autogen, 0)
	ec.defSubr3(nil, "insert-buffer-substring", ec.insertBufferSubstring_autogen, 1)
	ec.defSubr3(nil, "insert-byte", ec.insertByte_autogen, 2)
	ec.defSubr3(nil, "insert-char", ec.insertChar_autogen, 1)
	ec.defSubr1(nil, "internal--lock-narrowing", ec.internalLockNarrowing_autogen, 1)
	ec.defSubr1(nil, "internal--unlock-narrowing", ec.internalUnlockNarrowing_autogen, 1)
	ec.defSubr1(nil, "line-beginning-position", ec.lineBeginningPosition_autogen, 0)
	ec.defSubr1(nil, "line-end-position", ec.lineEndPosition_autogen, 0)
	ec.defSubr0(nil, "mark-marker", ec.markMarker_autogen)
	ec.defSubrM(nil, "message", ec.message_autogen, 1)
	ec.defSubrM(nil, "message-box", ec.messageBox_autogen, 1)
	ec.defSubrM(nil, "message-or-box", ec.messageOrBox_autogen, 1)
	ec.defSubr2(nil, "narrow-to-region", ec.narrowToRegion_autogen, 2)
	ec.defSubr3(nil, "ngettext", ec.ngettext_autogen, 3)
	ec.defSubr0(nil, "point", ec.point_autogen)
	ec.defSubr0(nil, "point-marker", ec.pointMarker_autogen)
	ec.defSubr0(nil, "point-max", ec.pointMax_autogen)
	ec.defSubr0(nil, "point-max-marker", ec.pointMaxMarker_autogen)
	ec.defSubr0(nil, "point-min", ec.pointMin_autogen)
	ec.defSubr0(nil, "point-min-marker", ec.pointMinMarker_autogen)
	ec.defSubr1(nil, "pos-bol", ec.posBol_autogen, 0)
	ec.defSubr1(nil, "pos-eol", ec.posEol_autogen, 0)
	ec.defSubr1(nil, "position-bytes", ec.positionBytes_autogen, 1)
	ec.defSubr0(nil, "preceding-char", ec.previousChar_autogen)
	ec.defSubrM(nil, "propertize", ec.propertize_autogen, 1)
	ec.defSubr0(nil, "region-beginning", ec.regionBeginning_autogen)
	ec.defSubr0(nil, "region-end", ec.regionEnd_autogen)
	ec.defSubr3(nil, "replace-buffer-contents", ec.replaceBufferContents_autogen, 1)
	ec.defSubrU(nil, "save-current-buffer", ec.saveCurrentBuffer_autogen, 0)
	ec.defSubrU(nil, "save-excursion", ec.saveExcursion_autogen, 0)
	ec.defSubrU(nil, "save-restriction", ec.saveRestriction_autogen, 0)
	ec.defSubr1(nil, "string-to-char", ec.stringToChar_autogen, 1)
	ec.defSubr5(nil, "subst-char-in-region", ec.substCharInRegion_autogen, 4)
	ec.defSubr0(nil, "system-name", ec.systemName_autogen)
	ec.defSubr3(nil, "translate-region-internal", ec.translateRegionInternal_autogen, 3)
	ec.defSubr5(nil, "transpose-regions", ec.transposeRegions_autogen, 4)
	ec.defSubr1(nil, "user-full-name", ec.userFullName_autogen, 0)
	ec.defSubr1(nil, "user-login-name", ec.userLoginName_autogen, 0)
	ec.defSubr0(nil, "user-real-login-name", ec.userRealLoginName_autogen)
	ec.defSubr0(nil, "user-real-uid", ec.userRealUid_autogen)
	ec.defSubr0(nil, "user-uid", ec.userUid_autogen)
	ec.defSubr0(nil, "widen", ec.widen_autogen)
	ec.defSubr1(nil, "module-load", ec.moduleLoad_autogen, 1)
	ec.defSubr0(nil, "daemon-initialized", ec.daemonInitialized_autogen)
	ec.defSubr0(nil, "daemonp", ec.daemonp_autogen)
	ec.defSubr2(nil, "dump-emacs", ec.dumpEmacs_autogen, 2)
	ec.defSubr0(nil, "invocation-directory", ec.invocationDirectory_autogen)
	ec.defSubr0(nil, "invocation-name", ec.invocationName_autogen)
	ec.defSubr2(nil, "kill-emacs", ec.killEmacs_autogen, 0)
	ec.defSubrU(nil, "and", ec.and_autogen, 0)
	ec.defSubrM(nil, "apply", ec.apply_autogen, 1)
	ec.defSubr5(nil, "autoload", ec.autoload_autogen, 2)
	ec.defSubr3(nil, "autoload-do-load", ec.autoloadDoLoad_autogen, 1)
	ec.defSubr1(nil, "backtrace--frames-from-thread", ec.backtraceFramesFromThread_autogen, 1)
	ec.defSubr2(nil, "backtrace--locals", ec.backtraceLocals_autogen, 1)
	ec.defSubr2(nil, "backtrace-debug", ec.backtraceDebug_autogen, 2)
	ec.defSubr3(nil, "backtrace-eval", ec.backtraceEval_autogen, 2)
	ec.defSubr3(nil, "backtrace-frame--internal", ec.backtraceFrameInternal_autogen, 3)
	ec.defSubrU(nil, "catch", ec.catch_autogen, 1)
	ec.defSubr2(nil, "commandp", ec.commandp_autogen, 1)
	ec.defSubrU(nil, "cond", ec.cond_autogen, 0)
	ec.defSubrU(nil, "condition-case", ec.conditionCase_autogen, 2)
	ec.defSubr1(nil, "default-toplevel-value", ec.defaultToplevelValue_autogen, 1)
	ec.defSubrU(nil, "defconst", ec.defconst_autogen, 2)
	ec.defSubr3(nil, "defconst-1", ec.defconst1_autogen, 2)
	ec.defSubrU(nil, "defvar", ec.defvar_autogen, 1)
	ec.defSubr3(nil, "defvar-1", ec.defvar1_autogen, 2)
	ec.defSubr3(nil, "defvaralias", ec.defvaralias_autogen, 2)
	ec.defSubr2(nil, "eval", ec.eval_autogen, 1)
	ec.defSubr1(nil, "fetch-bytecode", ec.fetchBytecode_autogen, 1)
	ec.defSubr1(nil, "func-arity", ec.funcArity_autogen, 1)
	ec.defSubrM(nil, "funcall", ec.funcall_autogen, 1)
	ec.defSubr3(nil, "funcall-with-delayed-message", ec.funcallWithDelayedMessage_autogen, 3)
	ec.defSubrU(nil, "function", ec.function_autogen, 1)
	ec.defSubr1(nil, "functionp", ec.functionp_autogen, 1)
	ec.defSubrU(nil, "if", ec.if_autogen, 2)
	ec.defSubr2(nil, "internal--define-uninitialized-variable", ec.internalDefineUninitializedVariable_autogen, 1)
	ec.defSubr1(nil, "internal-make-var-non-special", ec.makeVarNonSpecial_autogen, 1)
	ec.defSubrU(nil, "let", ec.let_autogen, 1)
	ec.defSubrU(nil, "let*", ec.letX_autogen, 1)
	ec.defSubr2(nil, "macroexpand", ec.macroexpand_autogen, 1)
	ec.defSubr2(nil, "mapbacktrace", ec.mapbacktrace_autogen, 1)
	ec.defSubrU(nil, "or", ec.or_autogen, 0)
	ec.defSubrU(nil, "prog1", ec.prog1_autogen, 1)
	ec.defSubrU(nil, "progn", ec.progn_autogen, 0)
	ec.defSubrU(nil, "quote", ec.quote_autogen, 1)
	ec.defSubrM(nil, "run-hook-with-args", ec.runHookWithArgs_autogen, 1)
	ec.defSubrM(nil, "run-hook-with-args-until-failure", ec.runHookWithArgsUntilFailure_autogen, 1)
	ec.defSubrM(nil, "run-hook-with-args-until-success", ec.runHookWithArgsUntilSuccess_autogen, 1)
	ec.defSubrM(nil, "run-hook-wrapped", ec.runHookWrapped_autogen, 2)
	ec.defSubrM(nil, "run-hooks", ec.runHooks_autogen, 0)
	ec.defSubr2(nil, "set-default-toplevel-value", ec.setDefaultToplevelValue_autogen, 2)
	ec.defSubrU(nil, "setq", ec.setq_autogen, 0)
	ec.defSubr2(nil, "signal", ec.signal_autogen, 2)
	ec.defSubr1(nil, "special-variable-p", ec.specialVariableP_autogen, 1)
	ec.defSubr2(nil, "throw", ec.throw_autogen, 2)
	ec.defSubrU(nil, "unwind-protect", ec.unwindProtect_autogen, 1)
	ec.defSubrU(nil, "while", ec.while_autogen, 1)
	ec.defSubr2(nil, "access-file", ec.accessFile_autogen, 2)
	ec.defSubr3(nil, "add-name-to-file", ec.addNameToFile_autogen, 2)
	ec.defSubr2(nil, "car-less-than-car", ec.carLessThanCar_autogen, 2)
	ec.defSubr0(nil, "clear-buffer-auto-save-failure", ec.clearBufferAutoSaveFailure_autogen)
	ec.defSubr6(nil, "copy-file", ec.copyFile_autogen, 2)
	ec.defSubr0(nil, "default-file-modes", ec.defaultFileModes_autogen)
	ec.defSubr1(nil, "delete-directory-internal", ec.deleteDirectoryInternal_autogen, 1)
	ec.defSubr2(nil, "delete-file", ec.deleteFile_autogen, 1)
	ec.defSubr1(nil, "directory-file-name", ec.directoryFileName_autogen, 1)
	ec.defSubr1(nil, "directory-name-p", ec.directoryNameP_autogen, 1)
	ec.defSubr2(nil, "do-auto-save", ec.doAutoSave_autogen, 0)
	ec.defSubr2(nil, "expand-file-name", ec.expandFileName_autogen, 1)
	ec.defSubr1(nil, "file-accessible-directory-p", ec.fileAccessibleDirectoryP_autogen, 1)
	ec.defSubr1(nil, "file-acl", ec.fileAcl_autogen, 1)
	ec.defSubr1(nil, "file-directory-p", ec.fileDirectoryP_autogen, 1)
	ec.defSubr1(nil, "file-executable-p", ec.fileExecutableP_autogen, 1)
	ec.defSubr1(nil, "file-exists-p", ec.fileExistsP_autogen, 1)
	ec.defSubr2(nil, "file-modes", ec.fileModes_autogen, 1)
	ec.defSubr1(nil, "file-name-absolute-p", ec.fileNameAbsoluteP_autogen, 1)
	ec.defSubr1(nil, "file-name-as-directory", ec.fileNameAsDirectory_autogen, 1)
	ec.defSubr1(nil, "file-name-case-insensitive-p", ec.fileNameCaseInsensitiveP_autogen, 1)
	ec.defSubrM(nil, "file-name-concat", ec.fileNameConcat_autogen, 1)
	ec.defSubr1(nil, "file-name-directory", ec.fileNameDirectory_autogen, 1)
	ec.defSubr1(nil, "file-name-nondirectory", ec.fileNameNondirectory_autogen, 1)
	ec.defSubr2(nil, "file-newer-than-file-p", ec.fileNewerThanFileP_autogen, 2)
	ec.defSubr1(nil, "file-readable-p", ec.fileReadableP_autogen, 1)
	ec.defSubr1(nil, "file-regular-p", ec.fileRegularP_autogen, 1)
	ec.defSubr1(nil, "file-selinux-context", ec.fileSelinuxContext_autogen, 1)
	ec.defSubr1(nil, "file-symlink-p", ec.fileSymlinkP_autogen, 1)
	ec.defSubr1(nil, "file-writable-p", ec.fileWritableP_autogen, 1)
	ec.defSubr2(nil, "find-file-name-handler", ec.findFileNameHandler_autogen, 2)
	ec.defSubr5(nil, "insert-file-contents", ec.insertFileContents_autogen, 1)
	ec.defSubr1(nil, "make-directory-internal", ec.makeDirectoryInternal_autogen, 1)
	ec.defSubr3(nil, "make-symbolic-link", ec.makeSymbolicLink_autogen, 2)
	ec.defSubr4(nil, "make-temp-file-internal", ec.makeTempFileInternal_autogen, 4)
	ec.defSubr1(nil, "make-temp-name", ec.makeTempName_autogen, 1)
	ec.defSubr0(nil, "next-read-file-uses-dialog-p", ec.nextReadFileUsesDialogP_autogen)
	ec.defSubr0(nil, "recent-auto-save-p", ec.recentAutoSaveP_autogen)
	ec.defSubr3(nil, "rename-file", ec.renameFile_autogen, 2)
	ec.defSubr2(nil, "set-binary-mode", ec.setBinaryMode_autogen, 2)
	ec.defSubr0(nil, "set-buffer-auto-saved", ec.setBufferAutoSaved_autogen)
	ec.defSubr1(nil, "set-default-file-modes", ec.setDefaultFileModes_autogen, 1)
	ec.defSubr2(nil, "set-file-acl", ec.setFileAcl_autogen, 2)
	ec.defSubr3(nil, "set-file-modes", ec.setFileModes_autogen, 2)
	ec.defSubr2(nil, "set-file-selinux-context", ec.setFileSelinuxContext_autogen, 2)
	ec.defSubr3(nil, "set-file-times", ec.setFileTimes_autogen, 1)
	ec.defSubr1(nil, "set-visited-file-modtime", ec.setVisitedFileModtime_autogen, 0)
	ec.defSubr1(nil, "substitute-in-file-name", ec.substituteInFileName_autogen, 1)
	ec.defSubr1(nil, "unhandled-file-name-directory", ec.unhandledFileNameDirectory_autogen, 1)
	ec.defSubr0(nil, "unix-sync", ec.unixSync_autogen)
	ec.defSubr1(nil, "verify-visited-file-modtime", ec.verifyVisitedFileModtime_autogen, 0)
	ec.defSubr0(nil, "visited-file-modtime", ec.visitedFileModtime_autogen)
	ec.defSubr7(nil, "write-region", ec.writeRegion_autogen, 3)
	ec.defSubr1(nil, "file-locked-p", ec.fileLockedP_autogen, 1)
	ec.defSubr1(nil, "lock-buffer", ec.lockBuffer_autogen, 0)
	ec.defSubr1(nil, "lock-file", ec.lockFile_autogen, 1)
	ec.defSubr0(nil, "unlock-buffer", ec.unlockBuffer_autogen)
	ec.defSubr1(nil, "unlock-file", ec.unlockFile_autogen, 1)
	ec.defSubr1(nil, "abs", ec.abs_autogen, 1)
	ec.defSubr1(nil, "acos", ec.acos_autogen, 1)
	ec.defSubr1(nil, "asin", ec.asin_autogen, 1)
	ec.defSubr2(nil, "atan", ec.atan_autogen, 1)
	ec.defSubr2(nil, "ceiling", ec.ceiling_autogen, 1)
	ec.defSubr2(nil, "copysign", ec.copysign_autogen, 2)
	ec.defSubr1(nil, "cos", ec.cos_autogen, 1)
	ec.defSubr1(nil, "exp", ec.exp_autogen, 1)
	ec.defSubr2(nil, "expt", ec.expt_autogen, 2)
	ec.defSubr1(nil, "fceiling", ec.fceiling_autogen, 1)
	ec.defSubr1(nil, "ffloor", ec.ffloor_autogen, 1)
	ec.defSubr1(nil, "float", ec.float_autogen, 1)
	ec.defSubr2(nil, "floor", ec.floor_autogen, 1)
	ec.defSubr1(nil, "frexp", ec.frexp_autogen, 1)
	ec.defSubr1(nil, "fround", ec.fround_autogen, 1)
	ec.defSubr1(nil, "ftruncate", ec.ftruncate_autogen, 1)
	ec.defSubr1(nil, "isnan", ec.isnan_autogen, 1)
	ec.defSubr2(nil, "ldexp", ec.ldexp_autogen, 2)
	ec.defSubr2(nil, "log", ec.log_autogen, 1)
	ec.defSubr1(nil, "logb", ec.logb_autogen, 1)
	ec.defSubr2(nil, "round", ec.round_autogen, 1)
	ec.defSubr1(nil, "sin", ec.sin_autogen, 1)
	ec.defSubr1(nil, "sqrt", ec.sqrt_autogen, 1)
	ec.defSubr1(nil, "tan", ec.tan_autogen, 1)
	ec.defSubr2(nil, "truncate", ec.truncate_autogen, 1)
	ec.defSubrM(nil, "append", ec.append_autogen, 0)
	ec.defSubr3(nil, "assoc", ec.assoc_autogen, 2)
	ec.defSubr2(nil, "assq", ec.assq_autogen, 2)
	ec.defSubr4(nil, "base64-decode-region", ec.base64DecodeRegion_autogen, 2)
	ec.defSubr3(nil, "base64-decode-string", ec.base64DecodeString_autogen, 1)
	ec.defSubr3(nil, "base64-encode-region", ec.base64EncodeRegion_autogen, 2)
	ec.defSubr2(nil, "base64-encode-string", ec.base64EncodeString_autogen, 1)
	ec.defSubr3(nil, "base64url-encode-region", ec.base64urlEncodeRegion_autogen, 2)
	ec.defSubr2(nil, "base64url-encode-string", ec.base64urlEncodeString_autogen, 1)
	ec.defSubr1(nil, "buffer-hash", ec.bufferHash_autogen, 0)
	ec.defSubr1(nil, "buffer-line-statistics", ec.bufferLineStatistics_autogen, 0)
	ec.defSubr1(nil, "clear-string", ec.clearString_autogen, 1)
	ec.defSubr1(nil, "clrhash", ec.clrhash_autogen, 1)
	ec.defSubr7(nil, "compare-strings", ec.compareStrings_autogen, 6)
	ec.defSubrM(nil, "concat", ec.concat_autogen, 0)
	ec.defSubr1(nil, "copy-alist", ec.copyAlist_autogen, 1)
	ec.defSubr1(nil, "copy-hash-table", ec.copyHashTable_autogen, 1)
	ec.defSubr1(nil, "copy-sequence", ec.copySequence_autogen, 1)
	ec.defSubr3(nil, "define-hash-table-test", ec.defineHashTableTest_autogen, 3)
	ec.defSubr2(nil, "delete", ec.delete_autogen, 2)
	ec.defSubr2(nil, "delq", ec.delq_autogen, 2)
	ec.defSubr2(nil, "elt", ec.elt_autogen, 2)
	ec.defSubr2(nil, "eql", ec.eql_autogen, 2)
	ec.defSubr2(nil, "equal", ec.equal_autogen, 2)
	ec.defSubr2(nil, "equal-including-properties", ec.equalIncludingProperties_autogen, 2)
	ec.defSubr2(nil, "featurep", ec.featurep_autogen, 1)
	ec.defSubr2(nil, "fillarray", ec.fillarray_autogen, 2)
	ec.defSubr2(nil, "get", ec.get_autogen, 2)
	ec.defSubr3(nil, "gethash", ec.gethash_autogen, 2)
	ec.defSubr1(nil, "hash-table-count", ec.hashTableCount_autogen, 1)
	ec.defSubr1(nil, "hash-table-p", ec.hashTableP_autogen, 1)
	ec.defSubr1(nil, "hash-table-rehash-size", ec.hashTableRehashSize_autogen, 1)
	ec.defSubr1(nil, "hash-table-rehash-threshold", ec.hashTableRehashThreshold_autogen, 1)
	ec.defSubr1(nil, "hash-table-size", ec.hashTableSize_autogen, 1)
	ec.defSubr1(nil, "hash-table-test", ec.hashTableTest_autogen, 1)
	ec.defSubr1(nil, "hash-table-weakness", ec.hashTableWeakness_autogen, 1)
	ec.defSubr1(nil, "identity", ec.identity_autogen, 1)
	ec.defSubr1(nil, "length", ec.length_autogen, 1)
	ec.defSubr2(nil, "length<", ec.lengthLess_autogen, 2)
	ec.defSubr2(nil, "length=", ec.lengthEqual_autogen, 2)
	ec.defSubr2(nil, "length>", ec.lengthGreater_autogen, 2)
	ec.defSubr2(nil, "line-number-at-pos", ec.lineNumberAtPos_autogen, 0)
	ec.defSubr1(nil, "load-average", ec.loadAverage_autogen, 0)
	ec.defSubr1(nil, "locale-info", ec.localeInfo_autogen, 1)
	ec.defSubrM(nil, "make-hash-table", ec.makeHashTable_autogen, 0)
	ec.defSubr2(nil, "mapc", ec.mapc_autogen, 2)
	ec.defSubr2(nil, "mapcan", ec.mapcan_autogen, 2)
	ec.defSubr2(nil, "mapcar", ec.mapcar_autogen, 2)
	ec.defSubr3(nil, "mapconcat", ec.mapconcat_autogen, 2)
	ec.defSubr2(nil, "maphash", ec.maphash_autogen, 2)
	ec.defSubr5(nil, "md5", ec.md5_autogen, 1)
	ec.defSubr2(nil, "member", ec.member_autogen, 2)
	ec.defSubr2(nil, "memq", ec.memq_autogen, 2)
	ec.defSubr2(nil, "memql", ec.memql_autogen, 2)
	ec.defSubrM(nil, "nconc", ec.nconc_autogen, 0)
	ec.defSubr1(nil, "nreverse", ec.nreverse_autogen, 1)
	ec.defSubr2(nil, "ntake", ec.ntake_autogen, 2)
	ec.defSubr2(nil, "nth", ec.nth_autogen, 2)
	ec.defSubr2(nil, "nthcdr", ec.nthcdr_autogen, 2)
	ec.defSubr1(nil, "object-intervals", ec.objectIntervals_autogen, 1)
	ec.defSubr3(nil, "plist-get", ec.plistGet_autogen, 2)
	ec.defSubr3(nil, "plist-member", ec.plistMember_autogen, 2)
	ec.defSubr4(nil, "plist-put", ec.plistPut_autogen, 3)
	ec.defSubr1(nil, "proper-list-p", ec.properListP_autogen, 1)
	ec.defSubr2(nil, "provide", ec.provide_autogen, 1)
	ec.defSubr3(nil, "put", ec.put_autogen, 3)
	ec.defSubr3(nil, "puthash", ec.puthash_autogen, 3)
	ec.defSubr1(nil, "random", ec.random_autogen, 0)
	ec.defSubr2(nil, "rassoc", ec.rassoc_autogen, 2)
	ec.defSubr2(nil, "rassq", ec.rassq_autogen, 2)
	ec.defSubr2(nil, "remhash", ec.remhash_autogen, 2)
	ec.defSubr3(nil, "require", ec.require_autogen, 1)
	ec.defSubr1(nil, "reverse", ec.reverse_autogen, 1)
	ec.defSubr1(nil, "safe-length", ec.safeLength_autogen, 1)
	ec.defSubr5(nil, "secure-hash", ec.secureHash_autogen, 2)
	ec.defSubr0(nil, "secure-hash-algorithms", ec.secureHashAlgorithms_autogen)
	ec.defSubr2(nil, "sort", ec.sort_autogen, 2)
	ec.defSubr1(nil, "string-as-multibyte", ec.stringAsMultibyte_autogen, 1)
	ec.defSubr1(nil, "string-as-unibyte", ec.stringAsUnibyte_autogen, 1)
	ec.defSubr1(nil, "string-bytes", ec.stringBytes_autogen, 1)
	ec.defSubr4(nil, "string-collate-equalp", ec.stringCollateEqualp_autogen, 2)
	ec.defSubr4(nil, "string-collate-lessp", ec.stringCollateLessp_autogen, 2)
	ec.defSubr3(nil, "string-distance", ec.stringDistance_autogen, 2)
	ec.defSubr2(nil, "string-equal", ec.stringEqual_autogen, 2)
	ec.defSubr2(nil, "string-lessp", ec.stringLessp_autogen, 2)
	ec.defSubr1(nil, "string-make-multibyte", ec.stringMakeMultibyte_autogen, 1)
	ec.defSubr1(nil, "string-make-unibyte", ec.stringMakeUnibyte_autogen, 1)
	ec.defSubr3(nil, "string-search", ec.stringSearch_autogen, 2)
	ec.defSubr1(nil, "string-to-multibyte", ec.stringToMultibyte_autogen, 1)
	ec.defSubr1(nil, "string-to-unibyte", ec.stringToUnibyte_autogen, 1)
	ec.defSubr2(nil, "string-version-lessp", ec.stringVersionLessp_autogen, 2)
	ec.defSubr3(nil, "substring", ec.substring_autogen, 1)
	ec.defSubr3(nil, "substring-no-properties", ec.substringNoProperties_autogen, 1)
	ec.defSubr1(nil, "sxhash-eq", ec.sxhashEq_autogen, 1)
	ec.defSubr1(nil, "sxhash-eql", ec.sxhashEql_autogen, 1)
	ec.defSubr1(nil, "sxhash-equal", ec.sxhashEqual_autogen, 1)
	ec.defSubr1(nil, "sxhash-equal-including-properties", ec.sxhashEqualIncludingProperties_autogen, 1)
	ec.defSubr2(nil, "take", ec.take_autogen, 2)
	ec.defSubrM(nil, "vconcat", ec.vconcat_autogen, 0)
	ec.defSubrM(nil, "widget-apply", ec.widgetApply_autogen, 2)
	ec.defSubr2(nil, "widget-get", ec.widgetGet_autogen, 2)
	ec.defSubr3(nil, "widget-put", ec.widgetPut_autogen, 3)
	ec.defSubr1(nil, "yes-or-no-p", ec.yesOrNoP_autogen, 1)
	ec.defSubr0(nil, "clear-font-cache", ec.clearFontCache_autogen)
	ec.defSubr2(nil, "close-font", ec.closeFont_autogen, 1)
	ec.defSubr2(nil, "draw-string", ec.drawString_autogen, 2)
	ec.defSubr2(nil, "find-font", ec.findFont_autogen, 1)
	ec.defSubr3(nil, "font-at", ec.fontAt_autogen, 1)
	ec.defSubr6(nil, "font-drive-otf", ec.fontDriveOtf_autogen, 6)
	ec.defSubr2(nil, "font-face-attributes", ec.fontFaceAttributes_autogen, 1)
	ec.defSubr1(nil, "font-family-list", ec.fontFamilyList_autogen, 0)
	ec.defSubr2(nil, "font-get", ec.fontGet_autogen, 2)
	ec.defSubr4(nil, "font-get-glyphs", ec.fontGetGlyphs_autogen, 3)
	ec.defSubr3(nil, "font-has-char-p", ec.fontHasCharP_autogen, 2)
	ec.defSubr2(nil, "font-info", ec.fontInfo_autogen, 1)
	ec.defSubr2(nil, "font-match-p", ec.fontMatchP_autogen, 2)
	ec.defSubr3(nil, "font-otf-alternates", ec.fontOtfAlternates_autogen, 3)
	ec.defSubr3(nil, "font-put", ec.fontPut_autogen, 3)
	ec.defSubr2(nil, "font-shape-gstring", ec.fontShapeGstring_autogen, 2)
	ec.defSubrM(nil, "font-spec", ec.fontSpec_autogen, 0)
	ec.defSubr2(nil, "font-variation-glyphs", ec.fontVariationGlyphs_autogen, 2)
	ec.defSubr2(nil, "font-xlfd-name", ec.fontXlfdName_autogen, 1)
	ec.defSubr2(nil, "fontp", ec.fontp_autogen, 1)
	ec.defSubr1(nil, "frame-font-cache", ec.frameFontCache_autogen, 0)
	ec.defSubr2(nil, "internal-char-font", ec.internalCharFont_autogen, 1)
	ec.defSubr4(nil, "list-fonts", ec.listFonts_autogen, 1)
	ec.defSubr3(nil, "open-font", ec.openFont_autogen, 1)
	ec.defSubr1(nil, "query-font", ec.queryFont_autogen, 1)
	ec.defSubr3(nil, "fontset-font", ec.fontsetFont_autogen, 2)
	ec.defSubr2(nil, "fontset-info", ec.fontsetInfo_autogen, 1)
	ec.defSubr0(nil, "fontset-list", ec.fontsetList_autogen)
	ec.defSubr0(nil, "fontset-list-all", ec.fontsetListAll_autogen)
	ec.defSubr2(nil, "new-fontset", ec.newFontset_autogen, 2)
	ec.defSubr2(nil, "query-fontset", ec.queryFontset_autogen, 1)
	ec.defSubr5(nil, "set-fontset-font", ec.setFontsetFont_autogen, 3)
	ec.defSubr2(nil, "delete-frame", ec.deleteFrame_autogen, 0)
	ec.defSubr2(nil, "frame--set-was-invisible", ec.frameSetWasInvisible_autogen, 2)
	ec.defSubr2(nil, "frame-after-make-frame", ec.frameAfterMakeFrame_autogen, 2)
	ec.defSubr2(nil, "frame-ancestor-p", ec.frameAncestorP_autogen, 2)
	ec.defSubr1(nil, "frame-bottom-divider-width", ec.bottomDividerWidth_autogen, 0)
	ec.defSubr1(nil, "frame-char-height", ec.frameCharHeight_autogen, 0)
	ec.defSubr1(nil, "frame-char-width", ec.frameCharWidth_autogen, 0)
	ec.defSubr1(nil, "frame-child-frame-border-width", ec.frameChildFrameBorderWidth_autogen, 0)
	ec.defSubr1(nil, "frame-focus", ec.frameFocus_autogen, 0)
	ec.defSubr1(nil, "frame-fringe-width", ec.fringeWidth_autogen, 0)
	ec.defSubr1(nil, "frame-internal-border-width", ec.frameInternalBorderWidth_autogen, 0)
	ec.defSubr0(nil, "frame-list", ec.frameList_autogen)
	ec.defSubr1(nil, "frame-live-p", ec.frameLiveP_autogen, 1)
	ec.defSubr1(nil, "frame-native-height", ec.frameNativeHeight_autogen, 0)
	ec.defSubr1(nil, "frame-native-width", ec.frameNativeWidth_autogen, 0)
	ec.defSubr2(nil, "frame-parameter", ec.frameParameter_autogen, 2)
	ec.defSubr1(nil, "frame-parameters", ec.frameParameters_autogen, 0)
	ec.defSubr1(nil, "frame-parent", ec.frameParent_autogen, 0)
	ec.defSubr1(nil, "frame-pointer-visible-p", ec.framePointerVisibleP_autogen, 0)
	ec.defSubr1(nil, "frame-position", ec.framePosition_autogen, 0)
	ec.defSubr1(nil, "frame-right-divider-width", ec.rightDividerWidth_autogen, 0)
	ec.defSubr1(nil, "frame-scale-factor", ec.frameScaleFactor_autogen, 0)
	ec.defSubr1(nil, "frame-scroll-bar-height", ec.scrollBarHeight_autogen, 0)
	ec.defSubr1(nil, "frame-scroll-bar-width", ec.scrollBarWidth_autogen, 0)
	ec.defSubr1(nil, "frame-text-cols", ec.frameTextCols_autogen, 0)
	ec.defSubr1(nil, "frame-text-height", ec.frameTextHeight_autogen, 0)
	ec.defSubr1(nil, "frame-text-lines", ec.frameTextLines_autogen, 0)
	ec.defSubr1(nil, "frame-text-width", ec.frameTextWidth_autogen, 0)
	ec.defSubr1(nil, "frame-total-cols", ec.frameTotalCols_autogen, 0)
	ec.defSubr1(nil, "frame-total-lines", ec.frameTotalLines_autogen, 0)
	ec.defSubr1(nil, "frame-visible-p", ec.frameVisibleP_autogen, 1)
	ec.defSubr1(nil, "frame-window-state-change", ec.frameWindowStateChange_autogen, 0)
	ec.defSubr4(nil, "frame-windows-min-size", ec.frameWindowsMinSize_autogen, 4)
	ec.defSubr1(nil, "framep", ec.framep_autogen, 1)
	ec.defSubr1(nil, "handle-switch-frame", ec.handleSwitchFrame_autogen, 1)
	ec.defSubr1(nil, "iconify-frame", ec.iconifyFrame_autogen, 0)
	ec.defSubr0(nil, "last-nonminibuffer-frame", ec.lastNonminibufFrame_autogen)
	ec.defSubr1(nil, "lower-frame", ec.lowerFrame_autogen, 0)
	ec.defSubr2(nil, "make-frame-invisible", ec.makeFrameInvisible_autogen, 0)
	ec.defSubr1(nil, "make-frame-visible", ec.makeFrameVisible_autogen, 0)
	ec.defSubr1(nil, "make-terminal-frame", ec.makeTerminalFrame_autogen, 1)
	ec.defSubr2(nil, "modify-frame-parameters", ec.modifyFrameParameters_autogen, 2)
	ec.defSubr0(nil, "mouse-pixel-position", ec.mousePixelPosition_autogen)
	ec.defSubr0(nil, "mouse-position", ec.mousePosition_autogen)
	ec.defSubr2(nil, "next-frame", ec.nextFrame_autogen, 0)
	ec.defSubr0(nil, "old-selected-frame", ec.oldSelectedFrame_autogen)
	ec.defSubr2(nil, "previous-frame", ec.previousFrame_autogen, 0)
	ec.defSubr1(nil, "raise-frame", ec.raiseFrame_autogen, 0)
	ec.defSubr1(nil, "reconsider-frame-fonts", ec.reconsiderFrameFonts_autogen, 1)
	ec.defSubr2(nil, "redirect-frame-focus", ec.redirectFrameFocus_autogen, 1)
	ec.defSubr2(nil, "select-frame", ec.selectFrame_autogen, 1)
	ec.defSubr0(nil, "selected-frame", ec.selectedFrame_autogen)
	ec.defSubr4(nil, "set-frame-height", ec.setFrameHeight_autogen, 2)
	ec.defSubr3(nil, "set-frame-position", ec.setFramePosition_autogen, 3)
	ec.defSubr4(nil, "set-frame-size", ec.setFrameSize_autogen, 3)
	ec.defSubr4(nil, "set-frame-width", ec.setFrameWidth_autogen, 2)
	ec.defSubr2(nil, "set-frame-window-state-change", ec.setFrameWindowStateChange_autogen, 0)
	ec.defSubr3(nil, "set-mouse-pixel-position", ec.setMousePixelPosition_autogen, 3)
	ec.defSubr3(nil, "set-mouse-position", ec.setMousePosition_autogen, 3)
	ec.defSubr1(nil, "tool-bar-pixel-width", ec.toolBarPixelWidth_autogen, 0)
	ec.defSubr0(nil, "visible-frame-list", ec.visibleFrameList_autogen)
	ec.defSubr1(nil, "window-system", ec.windowSystem_autogen, 0)
	ec.defSubr2(nil, "x-focus-frame", ec.xFocusFrame_autogen, 1)
	ec.defSubr4(nil, "x-get-resource", ec.xGetResource_autogen, 2)
	ec.defSubr1(nil, "x-parse-geometry", ec.xParseGeometry_autogen, 1)
	ec.defSubr5(nil, "define-fringe-bitmap", ec.defineFringeBitmap_autogen, 2)
	ec.defSubr1(nil, "destroy-fringe-bitmap", ec.destroyFringeBitmap_autogen, 1)
	ec.defSubr2(nil, "fringe-bitmaps-at-pos", ec.fringeBitmapsAtPos_autogen, 0)
	ec.defSubr2(nil, "set-fringe-bitmap-face", ec.setFringeBitmapFace_autogen, 1)
	ec.defSubr3(nil, "gfile-add-watch", ec.gfileAddWatch_autogen, 3)
	ec.defSubr1(nil, "gfile-monitor-name", ec.gfileMonitorName_autogen, 1)
	ec.defSubr1(nil, "gfile-rm-watch", ec.gfileRmWatch_autogen, 1)
	ec.defSubr1(nil, "gfile-valid-p", ec.gfileValidP_autogen, 1)
	ec.defSubr2(nil, "gnutls-asynchronous-parameters", ec.gnutlsAsynchronousParameters_autogen, 2)
	ec.defSubr0(nil, "gnutls-available-p", ec.gnutlsAvailableP_autogen)
	ec.defSubr3(nil, "gnutls-boot", ec.gnutlsBoot_autogen, 3)
	ec.defSubr2(nil, "gnutls-bye", ec.gnutlsBye_autogen, 2)
	ec.defSubr0(nil, "gnutls-ciphers", ec.gnutlsCiphers_autogen)
	ec.defSubr1(nil, "gnutls-deinit", ec.gnutlsDeinit_autogen, 1)
	ec.defSubr0(nil, "gnutls-digests", ec.gnutlsDigests_autogen)
	ec.defSubr1(nil, "gnutls-error-fatalp", ec.gnutlsErrorFatalp_autogen, 1)
	ec.defSubr1(nil, "gnutls-error-string", ec.gnutlsErrorString_autogen, 1)
	ec.defSubr1(nil, "gnutls-errorp", ec.gnutlsErrorp_autogen, 1)
	ec.defSubr1(nil, "gnutls-format-certificate", ec.gnutlsFormatCertificate_autogen, 1)
	ec.defSubr1(nil, "gnutls-get-initstage", ec.gnutlsGetInitstage_autogen, 1)
	ec.defSubr2(nil, "gnutls-hash-digest", ec.gnutlsHashDigest_autogen, 2)
	ec.defSubr3(nil, "gnutls-hash-mac", ec.gnutlsHashMac_autogen, 3)
	ec.defSubr0(nil, "gnutls-macs", ec.gnutlsMacs_autogen)
	ec.defSubr1(nil, "gnutls-peer-status", ec.gnutlsPeerStatus_autogen, 1)
	ec.defSubr1(nil, "gnutls-peer-status-warning-describe", ec.gnutlsPeerStatusWarningDescribe_autogen, 1)
	ec.defSubr5(nil, "gnutls-symmetric-decrypt", ec.gnutlsSymmetricDecrypt_autogen, 4)
	ec.defSubr5(nil, "gnutls-symmetric-encrypt", ec.gnutlsSymmetricEncrypt_autogen, 4)
	ec.defSubr1(nil, "haiku-display-monitor-attributes-list", ec.haikuDisplayMonitorAttributesList_autogen, 0)
	ec.defSubr2(nil, "haiku-frame-edges", ec.haikuFrameEdges_autogen, 0)
	ec.defSubr1(nil, "haiku-frame-geometry", ec.haikuFrameGeometry_autogen, 0)
	ec.defSubr1(nil, "haiku-frame-list-z-order", ec.haikuFrameListZOrder_autogen, 0)
	ec.defSubr3(nil, "haiku-frame-restack", ec.haikuFrameRestack_autogen, 2)
	ec.defSubr0(nil, "haiku-get-version-string", ec.haikuGetVersionString_autogen)
	ec.defSubr0(nil, "haiku-mouse-absolute-pixel-position", ec.haikuMouseAbsolutePixelPosition_autogen)
	ec.defSubr2(nil, "haiku-put-resource", ec.haikuPutResource_autogen, 2)
	ec.defSubr6(nil, "haiku-read-file-name", ec.haikuReadFileName_autogen, 1)
	ec.defSubr1(nil, "haiku-save-session-reply", ec.haikuSaveSessionReply_autogen, 1)
	ec.defSubr2(nil, "haiku-set-mouse-absolute-pixel-position", ec.haikuSetMouseAbsolutePixelPosition_autogen, 2)
	ec.defSubr1(nil, "x-close-connection", ec.xCloseConnection_autogen, 1)
	ec.defSubr1(nil, "x-close-connection", ec.xCloseConnection_1_autogen, 1)
	ec.defSubr1(nil, "x-close-connection", ec.xCloseConnection_2_autogen, 1)
	ec.defSubr1(nil, "x-close-connection", ec.xCloseConnection_3_autogen, 1)
	ec.defSubr1(nil, "x-create-frame", ec.xCreateFrame_autogen, 1)
	ec.defSubr1(nil, "x-create-frame", ec.xCreateFrame_1_autogen, 1)
	ec.defSubr1(nil, "x-create-frame", ec.xCreateFrame_2_autogen, 1)
	ec.defSubr1(nil, "x-create-frame", ec.xCreateFrame_3_autogen, 1)
	ec.defSubr1(nil, "x-display-backing-store", ec.xDisplayBackingStore_autogen, 0)
	ec.defSubr1(nil, "x-display-backing-store", ec.xDisplayBackingStore_1_autogen, 0)
	ec.defSubr1(nil, "x-display-backing-store", ec.xDisplayBackingStore_2_autogen, 0)
	ec.defSubr1(nil, "x-display-backing-store", ec.xDisplayBackingStore_3_autogen, 0)
	ec.defSubr1(nil, "x-display-color-cells", ec.xDisplayColorCells_autogen, 0)
	ec.defSubr1(nil, "x-display-color-cells", ec.xDisplayColorCells_1_autogen, 0)
	ec.defSubr1(nil, "x-display-color-cells", ec.xDisplayColorCells_2_autogen, 0)
	ec.defSubr1(nil, "x-display-color-cells", ec.xDisplayColorCells_3_autogen, 0)
	ec.defSubr1(nil, "x-display-grayscale-p", ec.xDisplayGrayscaleP_autogen, 0)
	ec.defSubr1(nil, "x-display-grayscale-p", ec.xDisplayGrayscaleP_1_autogen, 0)
	ec.defSubr1(nil, "x-display-grayscale-p", ec.xDisplayGrayscaleP_2_autogen, 0)
	ec.defSubr1(nil, "x-display-grayscale-p", ec.xDisplayGrayscaleP_3_autogen, 0)
	ec.defSubr0(nil, "x-display-list", ec.xDisplayList_autogen)
	ec.defSubr0(nil, "x-display-list", ec.xDisplayList_1_autogen)
	ec.defSubr0(nil, "x-display-list", ec.xDisplayList_2_autogen)
	ec.defSubr0(nil, "x-display-list", ec.xDisplayList_3_autogen)
	ec.defSubr1(nil, "x-display-mm-height", ec.xDisplayMmHeight_autogen, 0)
	ec.defSubr1(nil, "x-display-mm-height", ec.xDisplayMmHeight_1_autogen, 0)
	ec.defSubr1(nil, "x-display-mm-height", ec.xDisplayMmHeight_2_autogen, 0)
	ec.defSubr1(nil, "x-display-mm-height", ec.xDisplayMmHeight_3_autogen, 0)
	ec.defSubr1(nil, "x-display-mm-width", ec.xDisplayMmWidth_autogen, 0)
	ec.defSubr1(nil, "x-display-mm-width", ec.xDisplayMmWidth_1_autogen, 0)
	ec.defSubr1(nil, "x-display-mm-width", ec.xDisplayMmWidth_2_autogen, 0)
	ec.defSubr1(nil, "x-display-mm-width", ec.xDisplayMmWidth_3_autogen, 0)
	ec.defSubr1(nil, "x-display-pixel-height", ec.xDisplayPixelHeight_autogen, 0)
	ec.defSubr1(nil, "x-display-pixel-height", ec.xDisplayPixelHeight_1_autogen, 0)
	ec.defSubr1(nil, "x-display-pixel-height", ec.xDisplayPixelHeight_2_autogen, 0)
	ec.defSubr1(nil, "x-display-pixel-height", ec.xDisplayPixelHeight_3_autogen, 0)
	ec.defSubr1(nil, "x-display-pixel-width", ec.xDisplayPixelWidth_autogen, 0)
	ec.defSubr1(nil, "x-display-pixel-width", ec.xDisplayPixelWidth_1_autogen, 0)
	ec.defSubr1(nil, "x-display-pixel-width", ec.xDisplayPixelWidth_2_autogen, 0)
	ec.defSubr1(nil, "x-display-pixel-width", ec.xDisplayPixelWidth_3_autogen, 0)
	ec.defSubr1(nil, "x-display-planes", ec.xDisplayPlanes_autogen, 0)
	ec.defSubr1(nil, "x-display-planes", ec.xDisplayPlanes_1_autogen, 0)
	ec.defSubr1(nil, "x-display-planes", ec.xDisplayPlanes_2_autogen, 0)
	ec.defSubr1(nil, "x-display-planes", ec.xDisplayPlanes_3_autogen, 0)
	ec.defSubr1(nil, "x-display-save-under", ec.xDisplaySaveUnder_autogen, 0)
	ec.defSubr1(nil, "x-display-save-under", ec.xDisplaySaveUnder_1_autogen, 0)
	ec.defSubr1(nil, "x-display-save-under", ec.xDisplaySaveUnder_2_autogen, 0)
	ec.defSubr1(nil, "x-display-save-under", ec.xDisplaySaveUnder_3_autogen, 0)
	ec.defSubr1(nil, "x-display-screens", ec.xDisplayScreens_autogen, 0)
	ec.defSubr1(nil, "x-display-screens", ec.xDisplayScreens_1_autogen, 0)
	ec.defSubr1(nil, "x-display-screens", ec.xDisplayScreens_2_autogen, 0)
	ec.defSubr1(nil, "x-display-screens", ec.xDisplayScreens_3_autogen, 0)
	ec.defSubr1(nil, "x-display-visual-class", ec.xDisplayVisualClass_autogen, 0)
	ec.defSubr1(nil, "x-display-visual-class", ec.xDisplayVisualClass_1_autogen, 0)
	ec.defSubr1(nil, "x-display-visual-class", ec.xDisplayVisualClass_2_autogen, 0)
	ec.defSubr1(nil, "x-display-visual-class", ec.xDisplayVisualClass_3_autogen, 0)
	ec.defSubr1(nil, "x-double-buffered-p", ec.xDoubleBufferedP_autogen, 0)
	ec.defSubr1(nil, "x-double-buffered-p", ec.xDoubleBufferedP_1_autogen, 0)
	ec.defSubr0(nil, "x-hide-tip", ec.xHideTip_autogen)
	ec.defSubr0(nil, "x-hide-tip", ec.xHideTip_1_autogen)
	ec.defSubr0(nil, "x-hide-tip", ec.xHideTip_2_autogen)
	ec.defSubr0(nil, "x-hide-tip", ec.xHideTip_3_autogen)
	ec.defSubr3(nil, "x-open-connection", ec.xOpenConnection_autogen, 1)
	ec.defSubr3(nil, "x-open-connection", ec.xOpenConnection_1_autogen, 1)
	ec.defSubr3(nil, "x-open-connection", ec.xOpenConnection_2_autogen, 1)
	ec.defSubr3(nil, "x-open-connection", ec.xOpenConnection_3_autogen, 1)
	ec.defSubr1(nil, "x-server-vendor", ec.xServerVendor_autogen, 0)
	ec.defSubr1(nil, "x-server-vendor", ec.xServerVendor_1_autogen, 0)
	ec.defSubr1(nil, "x-server-vendor", ec.xServerVendor_2_autogen, 0)
	ec.defSubr1(nil, "x-server-version", ec.xServerVersion_autogen, 0)
	ec.defSubr1(nil, "x-server-version", ec.xServerVersion_1_autogen, 0)
	ec.defSubr1(nil, "x-server-version", ec.xServerVersion_2_autogen, 0)
	ec.defSubr6(nil, "x-show-tip", ec.xShowTip_autogen, 1)
	ec.defSubr6(nil, "x-show-tip", ec.xShowTip_1_autogen, 1)
	ec.defSubr6(nil, "x-show-tip", ec.xShowTip_2_autogen, 1)
	ec.defSubr6(nil, "x-show-tip", ec.xShowTip_3_autogen, 1)
	ec.defSubr2(nil, "xw-color-defined-p", ec.xwColorDefinedP_autogen, 1)
	ec.defSubr2(nil, "xw-color-defined-p", ec.xwColorDefinedP_1_autogen, 1)
	ec.defSubr2(nil, "xw-color-defined-p", ec.xwColorDefinedP_2_autogen, 1)
	ec.defSubr2(nil, "xw-color-defined-p", ec.xwColorDefinedP_3_autogen, 1)
	ec.defSubr2(nil, "xw-color-values", ec.xwColorValues_autogen, 1)
	ec.defSubr2(nil, "xw-color-values", ec.xwColorValues_1_autogen, 1)
	ec.defSubr2(nil, "xw-color-values", ec.xwColorValues_2_autogen, 1)
	ec.defSubr2(nil, "xw-color-values", ec.xwColorValues_3_autogen, 1)
	ec.defSubr1(nil, "xw-display-color-p", ec.xwDisplayColorP_autogen, 0)
	ec.defSubr1(nil, "xw-display-color-p", ec.xwDisplayColorP_1_autogen, 0)
	ec.defSubr1(nil, "xw-display-color-p", ec.xwDisplayColorP_2_autogen, 0)
	ec.defSubr1(nil, "xw-display-color-p", ec.xwDisplayColorP_3_autogen, 0)
	ec.defSubr0(nil, "font-get-system-font", ec.fontGetSystemFont_autogen)
	ec.defSubr0(nil, "font-get-system-font", ec.fontGetSystemFont_1_autogen)
	ec.defSubr0(nil, "font-get-system-normal-font", ec.fontGetSystemNormalFont_autogen)
	ec.defSubr0(nil, "font-get-system-normal-font", ec.fontGetSystemNormalFont_1_autogen)
	ec.defSubr2(nil, "x-select-font", ec.xSelectFont_autogen, 0)
	ec.defSubr2(nil, "x-select-font", ec.xSelectFont_1_autogen, 0)
	ec.defSubr2(nil, "x-select-font", ec.xSelectFont_2_autogen, 0)
	ec.defSubr2(nil, "x-select-font", ec.xSelectFont_3_autogen, 0)
	ec.defSubr1(nil, "haiku-menu-bar-open", ec.haikuMenuBarOpen_autogen, 0)
	ec.defSubr0(nil, "menu-or-popup-active-p", ec.menuOrPopupActiveP_autogen)
	ec.defSubr0(nil, "menu-or-popup-active-p", ec.menuOrPopupActiveP_1_autogen)
	ec.defSubr0(nil, "menu-or-popup-active-p", ec.menuOrPopupActiveP_2_autogen)
	ec.defSubr0(nil, "menu-or-popup-active-p", ec.menuOrPopupActiveP_3_autogen)
	ec.defSubr4(nil, "haiku-drag-message", ec.haikuDragMessage_autogen, 2)
	ec.defSubr2(nil, "haiku-roster-launch", ec.haikuRosterLaunch_autogen, 2)
	ec.defSubr2(nil, "haiku-selection-data", ec.haikuSelectionData_autogen, 2)
	ec.defSubr1(nil, "haiku-selection-owner-p", ec.haikuSelectionOwnerP_autogen, 0)
	ec.defSubr4(nil, "haiku-selection-put", ec.haikuSelectionPut_autogen, 2)
	ec.defSubr1(nil, "haiku-selection-timestamp", ec.haikuSelectionTimestamp_autogen, 1)
	ec.defSubr2(nil, "haiku-send-message", ec.haikuSendMessage_autogen, 2)
	ec.defSubr3(nil, "haiku-write-node-attribute", ec.haikuWriteNodeAttribute_autogen, 3)
	ec.defSubr2(nil, "clear-image-cache", ec.clearImageCache_autogen, 0)
	ec.defSubr0(nil, "image-cache-size", ec.imageCacheSize_autogen)
	ec.defSubr2(nil, "image-flush", ec.imageFlush_autogen, 1)
	ec.defSubr2(nil, "image-mask-p", ec.imageMaskP_autogen, 1)
	ec.defSubr2(nil, "image-metadata", ec.imageMetadata_autogen, 1)
	ec.defSubr3(nil, "image-size", ec.imageSize_autogen, 1)
	ec.defSubr1(nil, "image-transforms-p", ec.imageTransformsP_autogen, 0)
	ec.defSubr0(nil, "imagemagick-types", ec.imagemagickTypes_autogen)
	ec.defSubr1(nil, "imagep", ec.imagep_autogen, 1)
	ec.defSubr1(nil, "init-image-library", ec.initImageLibrary_autogen, 1)
	ec.defSubr1(nil, "lookup-image", ec.lookupImage_autogen, 1)
	ec.defSubr7(nil, "compute-motion", ec.computeMotion_autogen, 7)
	ec.defSubr0(nil, "current-column", ec.currentColumn_autogen)
	ec.defSubr0(nil, "current-indentation", ec.currentIndentation_autogen)
	ec.defSubr2(nil, "indent-to", ec.indentTo_autogen, 1)
	ec.defSubr1(nil, "line-number-display-width", ec.lineNumberDisplayWidth_autogen, 0)
	ec.defSubr2(nil, "move-to-column", ec.moveToColumn_autogen, 1)
	ec.defSubr3(nil, "vertical-motion", ec.verticalMotion_autogen, 1)
	ec.defSubr3(nil, "inotify-add-watch", ec.inotifyAddWatch_autogen, 3)
	ec.defSubr0(nil, "inotify-allocated-p", ec.inotifyAllocatedP_autogen)
	ec.defSubr1(nil, "inotify-rm-watch", ec.inotifyRmWatch_autogen, 1)
	ec.defSubr1(nil, "inotify-valid-p", ec.inotifyValidP_autogen, 1)
	ec.defSubr0(nil, "inotify-watch-list", ec.inotifyWatchList_autogen)
	ec.defSubr0(nil, "combine-after-change-execute", ec.combineAfterChangeExecute_autogen)
	ec.defSubr0(nil, "json--available-p", ec.jsonAvailableP_autogen)
	ec.defSubrM(nil, "json-insert", ec.jsonInsert_autogen, 1)
	ec.defSubrM(nil, "json-parse-buffer", ec.jsonParseBuffer_autogen, 0)
	ec.defSubrM(nil, "json-parse-string", ec.jsonParseString_autogen, 1)
	ec.defSubrM(nil, "json-serialize", ec.jsonSerialize_autogen, 1)
	ec.defSubr0(nil, "abort-recursive-edit", ec.abortRecursiveEdit_autogen)
	ec.defSubr1(nil, "clear-this-command-keys", ec.clearThisCommandKeys_autogen, 0)
	ec.defSubr3(nil, "command-error-default-function", ec.commandErrorDefaultFunction_autogen, 3)
	ec.defSubr0(nil, "current-idle-time", ec.currentIdleTime_autogen)
	ec.defSubr0(nil, "current-input-mode", ec.currentInputMode_autogen)
	ec.defSubr0(nil, "discard-input", ec.discardInput_autogen)
	ec.defSubr1(nil, "event-convert-list", ec.eventConvertList_autogen, 1)
	ec.defSubr0(nil, "exit-recursive-edit", ec.exitRecursiveEdit_autogen)
	ec.defSubr1(nil, "input-pending-p", ec.inputPendingP_autogen, 0)
	ec.defSubr1(nil, "internal--track-mouse", ec.internalTrackMouse_autogen, 1)
	ec.defSubr1(nil, "internal-event-symbol-parse-modifiers", ec.eventSymbolParseModifiers_autogen, 1)
	ec.defSubr1(nil, "internal-handle-focus-in", ec.internalHandleFocusIn_autogen, 1)
	ec.defSubr1(nil, "lossage-size", ec.lossageSize_autogen, 0)
	ec.defSubr1(nil, "open-dribble-file", ec.openDribbleFile_autogen, 1)
	ec.defSubr2(nil, "posn-at-point", ec.posnAtPoint_autogen, 0)
	ec.defSubr4(nil, "posn-at-x-y", ec.posnAtXY_autogen, 2)
	ec.defSubr5(nil, "read-key-sequence", ec.readKeySequence_autogen, 1)
	ec.defSubr5(nil, "read-key-sequence-vector", ec.readKeySequenceVector_autogen, 1)
	ec.defSubr1(nil, "recent-keys", ec.recentKeys_autogen, 0)
	ec.defSubr0(nil, "recursion-depth", ec.recursionDepth_autogen)
	ec.defSubr0(nil, "recursive-edit", ec.recursiveEdit_autogen)
	ec.defSubr1(nil, "set--this-command-keys", ec.setThisCommandKeys_autogen, 1)
	ec.defSubr1(nil, "set-input-interrupt-mode", ec.setInputInterruptMode_autogen, 1)
	ec.defSubr2(nil, "set-input-meta-mode", ec.setInputMetaMode_autogen, 1)
	ec.defSubr4(nil, "set-input-mode", ec.setInputMode_autogen, 3)
	ec.defSubr2(nil, "set-output-flow-control", ec.setOutputFlowControl_autogen, 1)
	ec.defSubr1(nil, "set-quit-char", ec.setQuitChar_autogen, 1)
	ec.defSubr1(nil, "suspend-emacs", ec.suspendEmacs_autogen, 0)
	ec.defSubr0(nil, "this-command-keys", ec.thisCommandKeys_autogen)
	ec.defSubr0(nil, "this-command-keys-vector", ec.thisCommandKeysVector_autogen)
	ec.defSubr0(nil, "this-single-command-keys", ec.thisSingleCommandKeys_autogen)
	ec.defSubr0(nil, "this-single-command-raw-keys", ec.thisSingleCommandRawKeys_autogen)
	ec.defSubr0(nil, "top-level", ec.topLevel_autogen)
	ec.defSubr2(nil, "accessible-keymaps", ec.accessibleKeymaps_autogen, 1)
	ec.defSubr3(nil, "command-remapping", ec.commandRemapping_autogen, 1)
	ec.defSubr1(nil, "copy-keymap", ec.copyKeymap_autogen, 1)
	ec.defSubr2(nil, "current-active-maps", ec.currentActiveMaps_autogen, 0)
	ec.defSubr0(nil, "current-global-map", ec.currentGlobalMap_autogen)
	ec.defSubr0(nil, "current-local-map", ec.currentLocalMap_autogen)
	ec.defSubr0(nil, "current-minor-mode-maps", ec.currentMinorModeMaps_autogen)
	ec.defSubr4(nil, "define-key", ec.defineKey_autogen, 3)
	ec.defSubr3(nil, "describe-buffer-bindings", ec.describeBufferBindings_autogen, 1)
	ec.defSubr2(nil, "describe-vector", ec.describeVector_autogen, 1)
	ec.defSubr7(nil, "help--describe-vector", ec.helpDescribeVector_autogen, 7)
	ec.defSubr4(nil, "key-binding", ec.keyBinding_autogen, 1)
	ec.defSubr2(nil, "key-description", ec.keyDescription_autogen, 1)
	ec.defSubr2(nil, "keymap--get-keyelt", ec.keymapGetKeyelt_autogen, 2)
	ec.defSubr1(nil, "keymap-parent", ec.keymapParent_autogen, 1)
	ec.defSubr1(nil, "keymap-prompt", ec.keymapPrompt_autogen, 1)
	ec.defSubr1(nil, "keymapp", ec.keymapp_autogen, 1)
	ec.defSubr3(nil, "lookup-key", ec.lookupKey_autogen, 2)
	ec.defSubr1(nil, "make-keymap", ec.makeKeymap_autogen, 0)
	ec.defSubr1(nil, "make-sparse-keymap", ec.makeSparseKeymap_autogen, 0)
	ec.defSubr3(nil, "map-keymap", ec.mapKeymap_autogen, 2)
	ec.defSubr2(nil, "map-keymap-internal", ec.mapKeymapInternal_autogen, 2)
	ec.defSubr2(nil, "minor-mode-key-binding", ec.minorModeKeyBinding_autogen, 1)
	ec.defSubr2(nil, "set-keymap-parent", ec.setKeymapParent_autogen, 2)
	ec.defSubr2(nil, "single-key-description", ec.singleKeyDescription_autogen, 1)
	ec.defSubr1(nil, "text-char-description", ec.textCharDescription_autogen, 1)
	ec.defSubr1(nil, "use-global-map", ec.useGlobalMap_autogen, 1)
	ec.defSubr1(nil, "use-local-map", ec.useLocalMap_autogen, 1)
	ec.defSubr5(nil, "where-is-internal", ec.whereIsInternal_autogen, 1)
	ec.defSubr3(nil, "kqueue-add-watch", ec.kqueueAddWatch_autogen, 3)
	ec.defSubr1(nil, "kqueue-rm-watch", ec.kqueueRmWatch_autogen, 1)
	ec.defSubr1(nil, "kqueue-valid-p", ec.kqueueValidP_autogen, 1)
	ec.defSubr4(nil, "lcms-cam02-ucs", ec.lcmsCam02Ucs_autogen, 2)
	ec.defSubr5(nil, "lcms-cie-de2000", ec.lcmsCieDe2000_autogen, 2)
	ec.defSubr3(nil, "lcms-jab->jch", ec.lcmsJabToJch_autogen, 1)
	ec.defSubr3(nil, "lcms-jch->jab", ec.lcmsJchToJab_autogen, 1)
	ec.defSubr3(nil, "lcms-jch->xyz", ec.lcmsJchToXyz_autogen, 1)
	ec.defSubr1(nil, "lcms-temp->white-point", ec.lcmsTempToWhitePoint_autogen, 1)
	ec.defSubr3(nil, "lcms-xyz->jch", ec.lcmsXyzToJch_autogen, 1)
	ec.defSubr0(nil, "lcms2-available-p", ec.lcms2AvailableP_autogen)
	ec.defSubr5(nil, "eval-buffer", ec.evalBuffer_autogen, 0)
	ec.defSubr4(nil, "eval-region", ec.evalRegion_autogen, 2)
	ec.defSubr0(nil, "get-file-char", ec.getFileChar_autogen)
	ec.defSubr0(nil, "get-load-suffixes", ec.getLoadSuffixes_autogen)
	ec.defSubr2(nil, "intern", ec.intern_autogen, 1)
	ec.defSubr2(nil, "intern-soft", ec.internSoft_autogen, 1)
	ec.defSubr5(nil, "load", ec.load_autogen, 1)
	ec.defSubr4(nil, "locate-file-internal", ec.locateFileInternal_autogen, 2)
	ec.defSubr3(nil, "lread--substitute-object-in-subtree", ec.lreadSubstituteObjectInSubtree_autogen, 3)
	ec.defSubr2(nil, "mapatoms", ec.mapatoms_autogen, 1)
	ec.defSubr1(nil, "read", ec.read_autogen, 0)
	ec.defSubr3(nil, "read-char", ec.readChar_autogen, 0)
	ec.defSubr3(nil, "read-char-exclusive", ec.readCharExclusive_autogen, 0)
	ec.defSubr3(nil, "read-event", ec.readEvent_autogen, 0)
	ec.defSubr3(nil, "read-from-string", ec.readFromString_autogen, 1)
	ec.defSubr1(nil, "read-positioning-symbols", ec.readPositioningSymbols_autogen, 0)
	ec.defSubr2(nil, "unintern", ec.unintern_autogen, 1)
	ec.defSubr2(nil, "call-last-kbd-macro", ec.callLastKbdMacro_autogen, 0)
	ec.defSubr0(nil, "cancel-kbd-macro-events", ec.cancelKbdMacroEvents_autogen)
	ec.defSubr2(nil, "end-kbd-macro", ec.endKbdMacro_autogen, 0)
	ec.defSubr3(nil, "execute-kbd-macro", ec.executeKbdMacro_autogen, 1)
	ec.defSubr2(nil, "start-kbd-macro", ec.startKbdMacro_autogen, 1)
	ec.defSubr1(nil, "store-kbd-macro-event", ec.storeKbdMacroEvent_autogen, 1)
	ec.defSubr2(nil, "copy-marker", ec.copyMarker_autogen, 0)
	ec.defSubr1(nil, "marker-buffer", ec.markerBuffer_autogen, 1)
	ec.defSubr1(nil, "marker-insertion-type", ec.markerInsertionType_autogen, 1)
	ec.defSubr1(nil, "marker-position", ec.markerPosition_autogen, 1)
	ec.defSubr3(nil, "set-marker", ec.setMarker_autogen, 2)
	ec.defSubr2(nil, "set-marker-insertion-type", ec.setMarkerInsertionType_autogen, 2)
	ec.defSubr3(nil, "menu-bar-menu-at-x-y", ec.menuBarMenuAtXY_autogen, 2)
	ec.defSubr3(nil, "x-popup-dialog", ec.xPopupDialog_autogen, 2)
	ec.defSubr2(nil, "x-popup-menu", ec.xPopupMenu_autogen, 2)
	ec.defSubr0(nil, "abort-minibuffers", ec.abortMinibuffers_autogen)
	ec.defSubr0(nil, "active-minibuffer-window", ec.activeMinibufferWindow_autogen)
	ec.defSubr4(nil, "all-completions", ec.allCompletions_autogen, 2)
	ec.defSubr3(nil, "assoc-string", ec.assocString_autogen, 2)
	ec.defSubr8(nil, "completing-read", ec.completingRead_autogen, 2)
	ec.defSubr1(nil, "innermost-minibuffer-p", ec.innermostMinibufferP_autogen, 0)
	ec.defSubr3(nil, "internal-complete-buffer", ec.internalCompleteBuffer_autogen, 3)
	ec.defSubr0(nil, "minibuffer-contents", ec.minibufferContents_autogen)
	ec.defSubr0(nil, "minibuffer-contents-no-properties", ec.minibufferContentsNoProperties_autogen)
	ec.defSubr0(nil, "minibuffer-depth", ec.minibufferDepth_autogen)
	ec.defSubr1(nil, "minibuffer-innermost-command-loop-p", ec.minibufferInnermostCommandLoopP_autogen, 0)
	ec.defSubr0(nil, "minibuffer-prompt", ec.minibufferPrompt_autogen)
	ec.defSubr0(nil, "minibuffer-prompt-end", ec.minibufferPromptEnd_autogen)
	ec.defSubr2(nil, "minibufferp", ec.minibufferp_autogen, 0)
	ec.defSubr4(nil, "read-buffer", ec.readBuffer_autogen, 1)
	ec.defSubr2(nil, "read-command", ec.readCommand_autogen, 1)
	ec.defSubr7(nil, "read-from-minibuffer", ec.readFromMinibuffer_autogen, 1)
	ec.defSubr1(nil, "read-function", ec.readFunction_autogen, 1)
	ec.defSubr5(nil, "read-string", ec.readString_autogen, 1)
	ec.defSubr2(nil, "read-variable", ec.readVariable_autogen, 1)
	ec.defSubr1(nil, "set-minibuffer-window", ec.setMinibufferWindow_autogen, 1)
	ec.defSubr3(nil, "test-completion", ec.testCompletion_autogen, 2)
	ec.defSubr3(nil, "try-completion", ec.tryCompletion_autogen, 2)
	ec.defSubr2(nil, "dump-emacs-portable", ec.dumpEmacsPortable_autogen, 1)
	ec.defSubr2(nil, "dump-emacs-portable--sort-predicate", ec.dumpEmacsPortableSortPredicate_autogen, 2)
	ec.defSubr2(nil, "dump-emacs-portable--sort-predicate-copied", ec.dumpEmacsPortableSortPredicateCopied_autogen, 2)
	ec.defSubr0(nil, "pdumper-stats", ec.pdumperStats_autogen)
	ec.defSubr1(nil, "pgtk-backend-display-class", ec.pgtkBackendDisplayClass_autogen, 0)
	ec.defSubr1(nil, "pgtk-display-monitor-attributes-list", ec.pgtkDisplayMonitorAttributesList_autogen, 0)
	ec.defSubr1(nil, "pgtk-font-name", ec.pgtkFontName_autogen, 1)
	ec.defSubr2(nil, "pgtk-frame-edges", ec.pgtkFrameEdges_autogen, 0)
	ec.defSubr1(nil, "pgtk-frame-geometry", ec.pgtkFrameGeometry_autogen, 0)
	ec.defSubr3(nil, "pgtk-frame-restack", ec.pgtkFrameRestack_autogen, 2)
	ec.defSubr0(nil, "pgtk-get-page-setup", ec.pgtkGetPageSetup_autogen)
	ec.defSubr0(nil, "pgtk-mouse-absolute-pixel-position", ec.pgtkMouseAbsolutePixelPosition_autogen)
	ec.defSubr0(nil, "pgtk-page-setup-dialog", ec.pgtkPageSetupDialog_autogen)
	ec.defSubr1(nil, "pgtk-print-frames-dialog", ec.pgtkPrintFramesDialog_autogen, 0)
	ec.defSubr2(nil, "pgtk-set-monitor-scale-factor", ec.pgtkSetMonitorScaleFactor_autogen, 2)
	ec.defSubr2(nil, "pgtk-set-mouse-absolute-pixel-position", ec.pgtkSetMouseAbsolutePixelPosition_autogen, 2)
	ec.defSubr2(nil, "pgtk-set-resource", ec.pgtkSetResource_autogen, 2)
	ec.defSubr2(nil, "x-export-frames", ec.xExportFrames_autogen, 0)
	ec.defSubr2(nil, "x-export-frames", ec.xExportFrames_1_autogen, 0)
	ec.defSubr5(nil, "x-file-dialog", ec.xFileDialog_autogen, 2)
	ec.defSubr5(nil, "x-file-dialog", ec.xFileDialog_1_autogen, 2)
	ec.defSubr5(nil, "x-file-dialog", ec.xFileDialog_2_autogen, 2)
	ec.defSubr5(nil, "x-file-dialog", ec.xFileDialog_3_autogen, 2)
	ec.defSubr1(nil, "x-gtk-debug", ec.xGtkDebug_autogen, 1)
	ec.defSubr1(nil, "x-gtk-debug", ec.xGtkDebug_1_autogen, 1)
	ec.defSubr1(nil, "x-server-max-request-size", ec.xServerMaxRequestSize_autogen, 0)
	ec.defSubr1(nil, "x-server-max-request-size", ec.xServerMaxRequestSize_1_autogen, 0)
	ec.defSubr1(nil, "x-server-max-request-size", ec.xServerMaxRequestSize_2_autogen, 0)
	ec.defSubr2(nil, "pgtk-use-im-context", ec.pgtkUseImContext_autogen, 1)
	ec.defSubr1(nil, "x-menu-bar-open-internal", ec.xMenuBarOpenInternal_autogen, 0)
	ec.defSubr1(nil, "x-menu-bar-open-internal", ec.xMenuBarOpenInternal_1_autogen, 0)
	ec.defSubr1(nil, "x-menu-bar-open-internal", ec.xMenuBarOpenInternal_2_autogen, 0)
	ec.defSubr3(nil, "pgtk-disown-selection-internal", ec.pgtkDisownSelectionInternal_autogen, 1)
	ec.defSubr3(nil, "pgtk-drop-finish", ec.pgtkDropFinish_autogen, 3)
	ec.defSubr4(nil, "pgtk-get-selection-internal", ec.pgtkGetSelectionInternal_autogen, 2)
	ec.defSubr3(nil, "pgtk-own-selection-internal", ec.pgtkOwnSelectionInternal_autogen, 2)
	ec.defSubr2(nil, "pgtk-register-dnd-targets", ec.pgtkRegisterDndTargets_autogen, 2)
	ec.defSubr2(nil, "pgtk-selection-exists-p", ec.pgtkSelectionExistsP_autogen, 0)
	ec.defSubr2(nil, "pgtk-selection-owner-p", ec.pgtkSelectionOwnerP_autogen, 0)
	ec.defSubr2(nil, "pgtk-update-drop-status", ec.pgtkUpdateDropStatus_autogen, 2)
	ec.defSubr1(nil, "error-message-string", ec.errorMessageString_autogen, 1)
	ec.defSubr1(nil, "external-debugging-output", ec.externalDebuggingOutput_autogen, 1)
	ec.defSubr0(nil, "flush-standard-output", ec.flushStandardOutput_autogen)
	ec.defSubr3(nil, "prin1", ec.prin1_autogen, 1)
	ec.defSubr3(nil, "prin1-to-string", ec.prin1ToString_autogen, 1)
	ec.defSubr2(nil, "princ", ec.princ_autogen, 1)
	ec.defSubr2(nil, "print", ec.print_autogen, 1)
	ec.defSubr1(nil, "print--preprocess", ec.printPreprocess_autogen, 1)
	ec.defSubr2(nil, "redirect-debugging-output", ec.redirectDebuggingOutput_autogen, 1)
	ec.defSubr2(nil, "terpri", ec.terpri_autogen, 0)
	ec.defSubr2(nil, "write-char", ec.writeChar_autogen, 1)
	ec.defSubr4(nil, "accept-process-output", ec.acceptProcessOutput_autogen, 0)
	ec.defSubr2(nil, "continue-process", ec.continueProcess_autogen, 0)
	ec.defSubr1(nil, "delete-process", ec.deleteProcess_autogen, 0)
	ec.defSubr2(nil, "format-network-address", ec.formatNetworkAddress_autogen, 1)
	ec.defSubr1(nil, "get-buffer-process", ec.getBufferProcess_autogen, 1)
	ec.defSubr1(nil, "get-process", ec.getProcess_autogen, 1)
	ec.defSubr2(nil, "internal-default-interrupt-process", ec.internalDefaultInterruptProcess_autogen, 0)
	ec.defSubr2(nil, "internal-default-process-filter", ec.internalDefaultProcessFilter_autogen, 2)
	ec.defSubr2(nil, "internal-default-process-sentinel", ec.internalDefaultProcessSentinel_autogen, 2)
	ec.defSubr3(nil, "internal-default-signal-process", ec.internalDefaultSignalProcess_autogen, 2)
	ec.defSubr2(nil, "interrupt-process", ec.interruptProcess_autogen, 0)
	ec.defSubr2(nil, "kill-process", ec.killProcess_autogen, 0)
	ec.defSubr0(nil, "list-system-processes", ec.listSystemProcesses_autogen)
	ec.defSubrM(nil, "make-network-process", ec.makeNetworkProcess_autogen, 0)
	ec.defSubrM(nil, "make-pipe-process", ec.makePipeProcess_autogen, 0)
	ec.defSubrM(nil, "make-process", ec.makeProcess_autogen, 0)
	ec.defSubrM(nil, "make-serial-process", ec.makeSerialProcess_autogen, 0)
	ec.defSubr1(nil, "network-interface-info", ec.networkInterfaceInfo_autogen, 1)
	ec.defSubr2(nil, "network-interface-list", ec.networkInterfaceList_autogen, 0)
	ec.defSubr3(nil, "network-lookup-address-info", ec.networkLookupAddressInfo_autogen, 1)
	ec.defSubr1(nil, "num-processors", ec.numProcessors_autogen, 0)
	ec.defSubr1(nil, "process-attributes", ec.processAttributes_autogen, 1)
	ec.defSubr1(nil, "process-buffer", ec.processBuffer_autogen, 1)
	ec.defSubr1(nil, "process-coding-system", ec.processCodingSystem_autogen, 1)
	ec.defSubr1(nil, "process-command", ec.processCommand_autogen, 1)
	ec.defSubr1(nil, "process-connection", ec.processConnection_autogen, 1)
	ec.defSubr3(nil, "process-contact", ec.processContact_autogen, 1)
	ec.defSubr1(nil, "process-datagram-address", ec.processDatagramAddress_autogen, 1)
	ec.defSubr1(nil, "process-exit-status", ec.processExitStatus_autogen, 1)
	ec.defSubr1(nil, "process-filter", ec.processFilter_autogen, 1)
	ec.defSubr1(nil, "process-id", ec.processId_autogen, 1)
	ec.defSubr1(nil, "process-inherit-coding-system-flag", ec.processInheritCodingSystemFlag_autogen, 1)
	ec.defSubr0(nil, "process-list", ec.processList_autogen)
	ec.defSubr1(nil, "process-mark", ec.processMark_autogen, 1)
	ec.defSubr1(nil, "process-name", ec.processName_autogen, 1)
	ec.defSubr1(nil, "process-plist", ec.processPlist_autogen, 1)
	ec.defSubr1(nil, "process-query-on-exit-flag", ec.processQueryOnExitFlag_autogen, 1)
	ec.defSubr1(nil, "process-running-child-p", ec.processRunningChildP_autogen, 0)
	ec.defSubr1(nil, "process-send-eof", ec.processSendEof_autogen, 0)
	ec.defSubr3(nil, "process-send-region", ec.processSendRegion_autogen, 3)
	ec.defSubr2(nil, "process-send-string", ec.processSendString_autogen, 2)
	ec.defSubr1(nil, "process-sentinel", ec.processSentinel_autogen, 1)
	ec.defSubr1(nil, "process-status", ec.processStatus_autogen, 1)
	ec.defSubr1(nil, "process-thread", ec.processThread_autogen, 1)
	ec.defSubr2(nil, "process-tty-name", ec.processTtyName_autogen, 1)
	ec.defSubr1(nil, "process-type", ec.processType_autogen, 1)
	ec.defSubr1(nil, "processp", ec.processp_autogen, 1)
	ec.defSubr2(nil, "quit-process", ec.quitProcess_autogen, 0)
	ec.defSubrM(nil, "serial-process-configure", ec.serialProcessConfigure_autogen, 0)
	ec.defSubr4(nil, "set-network-process-option", ec.setNetworkProcessOption_autogen, 3)
	ec.defSubr2(nil, "set-process-buffer", ec.setProcessBuffer_autogen, 2)
	ec.defSubr3(nil, "set-process-coding-system", ec.setProcessCodingSystem_autogen, 1)
	ec.defSubr2(nil, "set-process-datagram-address", ec.setProcessDatagramAddress_autogen, 2)
	ec.defSubr2(nil, "set-process-filter", ec.setProcessFilter_autogen, 2)
	ec.defSubr2(nil, "set-process-inherit-coding-system-flag", ec.setProcessInheritCodingSystemFlag_autogen, 2)
	ec.defSubr2(nil, "set-process-plist", ec.setProcessPlist_autogen, 2)
	ec.defSubr2(nil, "set-process-query-on-exit-flag", ec.setProcessQueryOnExitFlag_autogen, 2)
	ec.defSubr2(nil, "set-process-sentinel", ec.setProcessSentinel_autogen, 2)
	ec.defSubr2(nil, "set-process-thread", ec.setProcessThread_autogen, 2)
	ec.defSubr3(nil, "set-process-window-size", ec.setProcessWindowSize_autogen, 3)
	ec.defSubr0(nil, "signal-names", ec.signalNames_autogen)
	ec.defSubr3(nil, "signal-process", ec.signalProcess_autogen, 2)
	ec.defSubr2(nil, "stop-process", ec.stopProcess_autogen, 0)
	ec.defSubr0(nil, "waiting-for-user-input-p", ec.waitingForUserInputP_autogen)
	ec.defSubr2(nil, "function-equal", ec.functionEqual_autogen, 2)
	ec.defSubr0(nil, "profiler-cpu-log", ec.profilerCpuLog_autogen)
	ec.defSubr0(nil, "profiler-cpu-running-p", ec.profilerCpuRunningP_autogen)
	ec.defSubr1(nil, "profiler-cpu-start", ec.profilerCpuStart_autogen, 1)
	ec.defSubr0(nil, "profiler-cpu-stop", ec.profilerCpuStop_autogen)
	ec.defSubr0(nil, "profiler-memory-log", ec.profilerMemoryLog_autogen)
	ec.defSubr0(nil, "profiler-memory-running-p", ec.profilerMemoryRunningP_autogen)
	ec.defSubr0(nil, "profiler-memory-start", ec.profilerMemoryStart_autogen)
	ec.defSubr0(nil, "profiler-memory-stop", ec.profilerMemoryStop_autogen)
	ec.defSubr2(nil, "looking-at", ec.lookingAt_autogen, 1)
	ec.defSubr1(nil, "match-beginning", ec.matchBeginning_autogen, 1)
	ec.defSubr3(nil, "match-data", ec.matchData_autogen, 0)
	ec.defSubr1(nil, "match-data--translate", ec.matchDataTranslate_autogen, 1)
	ec.defSubr1(nil, "match-end", ec.matchEnd_autogen, 1)
	ec.defSubr1(nil, "newline-cache-check", ec.newlineCacheCheck_autogen, 0)
	ec.defSubr2(nil, "posix-looking-at", ec.posixLookingAt_autogen, 1)
	ec.defSubr4(nil, "posix-search-backward", ec.posixSearchBackward_autogen, 1)
	ec.defSubr4(nil, "posix-search-forward", ec.posixSearchForward_autogen, 1)
	ec.defSubr4(nil, "posix-string-match", ec.posixStringMatch_autogen, 2)
	ec.defSubr4(nil, "re-search-backward", ec.reSearchBackward_autogen, 1)
	ec.defSubr4(nil, "re-search-forward", ec.reSearchForward_autogen, 1)
	ec.defSubr1(nil, "regexp-quote", ec.regexpQuote_autogen, 1)
	ec.defSubr5(nil, "replace-match", ec.replaceMatch_autogen, 1)
	ec.defSubr4(nil, "search-backward", ec.searchBackward_autogen, 1)
	ec.defSubr4(nil, "search-forward", ec.searchForward_autogen, 1)
	ec.defSubr2(nil, "set-match-data", ec.setMatchData_autogen, 1)
	ec.defSubr4(nil, "string-match", ec.stringMatch_autogen, 2)
	ec.defSubr1(nil, "play-sound-internal", ec.playSoundInternal_autogen, 1)
	ec.defSubr0(nil, "sqlite-available-p", ec.sqliteAvailableP_autogen)
	ec.defSubr1(nil, "sqlite-close", ec.sqliteClose_autogen, 1)
	ec.defSubr1(nil, "sqlite-columns", ec.sqliteColumns_autogen, 1)
	ec.defSubr1(nil, "sqlite-commit", ec.sqliteCommit_autogen, 1)
	ec.defSubr3(nil, "sqlite-execute", ec.sqliteExecute_autogen, 2)
	ec.defSubr1(nil, "sqlite-finalize", ec.sqliteFinalize_autogen, 1)
	ec.defSubr2(nil, "sqlite-load-extension", ec.sqliteLoadExtension_autogen, 2)
	ec.defSubr1(nil, "sqlite-more-p", ec.sqliteMoreP_autogen, 1)
	ec.defSubr1(nil, "sqlite-next", ec.sqliteNext_autogen, 1)
	ec.defSubr1(nil, "sqlite-open", ec.sqliteOpen_autogen, 0)
	ec.defSubr2(nil, "sqlite-pragma", ec.sqlitePragma_autogen, 2)
	ec.defSubr1(nil, "sqlite-rollback", ec.sqliteRollback_autogen, 1)
	ec.defSubr4(nil, "sqlite-select", ec.sqliteSelect_autogen, 2)
	ec.defSubr1(nil, "sqlite-transaction", ec.sqliteTransaction_autogen, 1)
	ec.defSubr0(nil, "sqlite-version", ec.sqliteVersion_autogen)
	ec.defSubr1(nil, "sqlitep", ec.sqlitep_autogen, 1)
	ec.defSubr0(nil, "backward-prefix-chars", ec.backwardPrefixChars_autogen)
	ec.defSubr1(nil, "char-syntax", ec.charSyntax_autogen, 1)
	ec.defSubr1(nil, "copy-syntax-table", ec.copySyntaxTable_autogen, 0)
	ec.defSubr1(nil, "forward-comment", ec.forwardComment_autogen, 1)
	ec.defSubr1(nil, "forward-word", ec.forwardWord_autogen, 0)
	ec.defSubr1(nil, "internal-describe-syntax-value", ec.internalDescribeSyntaxValue_autogen, 1)
	ec.defSubr1(nil, "matching-paren", ec.matchingParen_autogen, 1)
	ec.defSubr3(nil, "modify-syntax-entry", ec.modifySyntaxEntry_autogen, 2)
	ec.defSubr6(nil, "parse-partial-sexp", ec.parsePartialSexp_autogen, 2)
	ec.defSubr3(nil, "scan-lists", ec.scanLists_autogen, 3)
	ec.defSubr2(nil, "scan-sexps", ec.scanSexps_autogen, 2)
	ec.defSubr1(nil, "set-syntax-table", ec.setSyntaxTable_autogen, 1)
	ec.defSubr2(nil, "skip-chars-backward", ec.skipCharsBackward_autogen, 1)
	ec.defSubr2(nil, "skip-chars-forward", ec.skipCharsForward_autogen, 1)
	ec.defSubr2(nil, "skip-syntax-backward", ec.skipSyntaxBackward_autogen, 1)
	ec.defSubr2(nil, "skip-syntax-forward", ec.skipSyntaxForward_autogen, 1)
	ec.defSubr0(nil, "standard-syntax-table", ec.standardSyntaxTable_autogen)
	ec.defSubr1(nil, "string-to-syntax", ec.stringToSyntax_autogen, 1)
	ec.defSubr1(nil, "syntax-class-to-char", ec.syntaxClassToChar_autogen, 1)
	ec.defSubr0(nil, "syntax-table", ec.syntaxTable_autogen)
	ec.defSubr1(nil, "syntax-table-p", ec.syntaxTableP_autogen, 1)
	ec.defSubr0(nil, "get-internal-run-time", ec.getInternalRunTime_autogen)
	ec.defSubr1(nil, "controlling-tty-p", ec.controllingTtyP_autogen, 0)
	ec.defSubr0(nil, "gpm-mouse-start", ec.gpmMouseStart_autogen)
	ec.defSubr0(nil, "gpm-mouse-stop", ec.gpmMouseStop_autogen)
	ec.defSubr1(nil, "resume-tty", ec.resumeTty_autogen, 0)
	ec.defSubr1(nil, "suspend-tty", ec.suspendTty_autogen, 0)
	ec.defSubr1(nil, "tty--output-buffer-size", ec.ttyOutputBufferSize_autogen, 0)
	ec.defSubr2(nil, "tty--set-output-buffer-size", ec.ttySetOutputBufferSize_autogen, 1)
	ec.defSubr1(nil, "tty-display-color-cells", ec.ttyDisplayColorCells_autogen, 0)
	ec.defSubr1(nil, "tty-display-color-p", ec.ttyDisplayColorP_autogen, 0)
	ec.defSubr1(nil, "tty-no-underline", ec.ttyNoUnderline_autogen, 0)
	ec.defSubr1(nil, "tty-top-frame", ec.ttyTopFrame_autogen, 0)
	ec.defSubr1(nil, "tty-type", ec.ttyType_autogen, 0)
	ec.defSubr2(nil, "delete-terminal", ec.deleteTerminal_autogen, 0)
	ec.defSubr1(nil, "frame-terminal", ec.frameTerminal_autogen, 0)
	ec.defSubr3(nil, "set-terminal-parameter", ec.setTerminalParameter_autogen, 3)
	ec.defSubr0(nil, "terminal-list", ec.terminalList_autogen)
	ec.defSubr1(nil, "terminal-live-p", ec.terminalLiveP_autogen, 1)
	ec.defSubr1(nil, "terminal-name", ec.terminalName_autogen, 0)
	ec.defSubr2(nil, "terminal-parameter", ec.terminalParameter_autogen, 2)
	ec.defSubr1(nil, "terminal-parameters", ec.terminalParameters_autogen, 0)
	ec.defSubr5(nil, "add-face-text-property", ec.addFaceTextProperty_autogen, 3)
	ec.defSubr4(nil, "add-text-properties", ec.addTextProperties_autogen, 3)
	ec.defSubr3(nil, "get-char-property", ec.getCharProperty_autogen, 2)
	ec.defSubr3(nil, "get-char-property-and-overlay", ec.getCharPropertyAndOverlay_autogen, 2)
	ec.defSubr3(nil, "get-text-property", ec.getTextProperty_autogen, 2)
	ec.defSubr2(nil, "next-char-property-change", ec.nextCharPropertyChange_autogen, 1)
	ec.defSubr3(nil, "next-property-change", ec.nextPropertyChange_autogen, 1)
	ec.defSubr4(nil, "next-single-char-property-change", ec.nextSingleCharPropertyChange_autogen, 2)
	ec.defSubr4(nil, "next-single-property-change", ec.nextSinglePropertyChange_autogen, 2)
	ec.defSubr2(nil, "previous-char-property-change", ec.previousCharPropertyChange_autogen, 1)
	ec.defSubr3(nil, "previous-property-change", ec.previousPropertyChange_autogen, 1)
	ec.defSubr4(nil, "previous-single-char-property-change", ec.previousSingleCharPropertyChange_autogen, 2)
	ec.defSubr4(nil, "previous-single-property-change", ec.previousSinglePropertyChange_autogen, 2)
	ec.defSubr5(nil, "put-text-property", ec.putTextProperty_autogen, 4)
	ec.defSubr4(nil, "remove-list-of-text-properties", ec.removeListOfTextProperties_autogen, 3)
	ec.defSubr4(nil, "remove-text-properties", ec.removeTextProperties_autogen, 3)
	ec.defSubr4(nil, "set-text-properties", ec.setTextProperties_autogen, 3)
	ec.defSubr2(nil, "text-properties-at", ec.textPropertiesAt_autogen, 1)
	ec.defSubr5(nil, "text-property-any", ec.textPropertyAny_autogen, 4)
	ec.defSubr5(nil, "text-property-not-all", ec.textPropertyNotAll_autogen, 4)
	ec.defSubr0(nil, "all-threads", ec.allThreads_autogen)
	ec.defSubr1(nil, "condition-mutex", ec.conditionMutex_autogen, 1)
	ec.defSubr1(nil, "condition-name", ec.conditionName_autogen, 1)
	ec.defSubr2(nil, "condition-notify", ec.conditionNotify_autogen, 1)
	ec.defSubr1(nil, "condition-wait", ec.conditionWait_autogen, 1)
	ec.defSubr0(nil, "current-thread", ec.currentThread_autogen)
	ec.defSubr2(nil, "make-condition-variable", ec.makeConditionVariable_autogen, 1)
	ec.defSubr1(nil, "make-mutex", ec.makeMutex_autogen, 0)
	ec.defSubr2(nil, "make-thread", ec.makeThread_autogen, 1)
	ec.defSubr1(nil, "mutex-lock", ec.mutexLock_autogen, 1)
	ec.defSubr1(nil, "mutex-name", ec.mutexName_autogen, 1)
	ec.defSubr1(nil, "mutex-unlock", ec.mutexUnlock_autogen, 1)
	ec.defSubr1(nil, "thread--blocker", ec.threadBlocker_autogen, 1)
	ec.defSubr1(nil, "thread-join", ec.threadJoin_autogen, 1)
	ec.defSubr1(nil, "thread-last-error", ec.threadLastError_autogen, 0)
	ec.defSubr1(nil, "thread-live-p", ec.threadLiveP_autogen, 1)
	ec.defSubr1(nil, "thread-name", ec.threadName_autogen, 1)
	ec.defSubr3(nil, "thread-signal", ec.threadSignal_autogen, 3)
	ec.defSubr0(nil, "thread-yield", ec.threadYield_autogen)
	ec.defSubr0(nil, "current-cpu-time", ec.currentCpuTime_autogen)
	ec.defSubr0(nil, "current-time", ec.currentTime_autogen)
	ec.defSubr2(nil, "current-time-string", ec.currentTimeString_autogen, 0)
	ec.defSubr2(nil, "current-time-zone", ec.currentTimeZone_autogen, 0)
	ec.defSubr3(nil, "decode-time", ec.decodeTime_autogen, 0)
	ec.defSubrM(nil, "encode-time", ec.encodeTime_autogen, 1)
	ec.defSubr1(nil, "float-time", ec.floatTime_autogen, 0)
	ec.defSubr3(nil, "format-time-string", ec.formatTimeString_autogen, 1)
	ec.defSubr1(nil, "set-time-zone-rule", ec.setTimeZoneRule_autogen, 1)
	ec.defSubr2(nil, "time-add", ec.timeAdd_autogen, 2)
	ec.defSubr2(nil, "time-convert", ec.timeConvert_autogen, 1)
	ec.defSubr2(nil, "time-equal-p", ec.timeEqualP_autogen, 2)
	ec.defSubr2(nil, "time-less-p", ec.timeLessP_autogen, 2)
	ec.defSubr2(nil, "time-subtract", ec.timeSubtract_autogen, 2)
	ec.defSubr0(nil, "treesit-available-p", ec.treesitAvailableP_autogen)
	ec.defSubr1(nil, "treesit-compiled-query-p", ec.treesitCompiledQueryP_autogen, 1)
	ec.defSubr4(nil, "treesit-induce-sparse-tree", ec.treesitInduceSparseTree_autogen, 2)
	ec.defSubr1(nil, "treesit-language-abi-version", ec.treesitLanguageAbiVersion_autogen, 0)
	ec.defSubr2(nil, "treesit-language-available-p", ec.treesitLanguageAvailableP_autogen, 1)
	ec.defSubr1(nil, "treesit-library-abi-version", ec.treesitLibraryAbiVersion_autogen, 0)
	ec.defSubr2(nil, "treesit-node-check", ec.treesitNodeCheck_autogen, 2)
	ec.defSubr3(nil, "treesit-node-child", ec.treesitNodeChild_autogen, 2)
	ec.defSubr2(nil, "treesit-node-child-by-field-name", ec.treesitNodeChildByFieldName_autogen, 2)
	ec.defSubr2(nil, "treesit-node-child-count", ec.treesitNodeChildCount_autogen, 1)
	ec.defSubr4(nil, "treesit-node-descendant-for-range", ec.treesitNodeDescendantForRange_autogen, 3)
	ec.defSubr1(nil, "treesit-node-end", ec.treesitNodeEnd_autogen, 1)
	ec.defSubr2(nil, "treesit-node-eq", ec.treesitNodeEq_autogen, 2)
	ec.defSubr2(nil, "treesit-node-field-name-for-child", ec.treesitNodeFieldNameForChild_autogen, 2)
	ec.defSubr3(nil, "treesit-node-first-child-for-pos", ec.treesitNodeFirstChildForPos_autogen, 2)
	ec.defSubr2(nil, "treesit-node-next-sibling", ec.treesitNodeNextSibling_autogen, 1)
	ec.defSubr1(nil, "treesit-node-p", ec.treesitNodeP_autogen, 1)
	ec.defSubr1(nil, "treesit-node-parent", ec.treesitNodeParent_autogen, 1)
	ec.defSubr1(nil, "treesit-node-parser", ec.treesitNodeParser_autogen, 1)
	ec.defSubr2(nil, "treesit-node-prev-sibling", ec.treesitNodePrevSibling_autogen, 1)
	ec.defSubr1(nil, "treesit-node-start", ec.treesitNodeStart_autogen, 1)
	ec.defSubr1(nil, "treesit-node-string", ec.treesitNodeString_autogen, 1)
	ec.defSubr1(nil, "treesit-node-type", ec.treesitNodeType_autogen, 1)
	ec.defSubr2(nil, "treesit-parser-add-notifier", ec.treesitParserAddNotifier_autogen, 2)
	ec.defSubr1(nil, "treesit-parser-buffer", ec.treesitParserBuffer_autogen, 1)
	ec.defSubr3(nil, "treesit-parser-create", ec.treesitParserCreate_autogen, 1)
	ec.defSubr1(nil, "treesit-parser-delete", ec.treesitParserDelete_autogen, 1)
	ec.defSubr1(nil, "treesit-parser-included-ranges", ec.treesitParserIncludedRanges_autogen, 1)
	ec.defSubr1(nil, "treesit-parser-language", ec.treesitParserLanguage_autogen, 1)
	ec.defSubr1(nil, "treesit-parser-list", ec.treesitParserList_autogen, 0)
	ec.defSubr1(nil, "treesit-parser-notifiers", ec.treesitParserNotifiers_autogen, 1)
	ec.defSubr1(nil, "treesit-parser-p", ec.treesitParserP_autogen, 1)
	ec.defSubr2(nil, "treesit-parser-remove-notifier", ec.treesitParserRemoveNotifier_autogen, 2)
	ec.defSubr1(nil, "treesit-parser-root-node", ec.treesitParserRootNode_autogen, 1)
	ec.defSubr2(nil, "treesit-parser-set-included-ranges", ec.treesitParserSetIncludedRanges_autogen, 2)
	ec.defSubr1(nil, "treesit-pattern-expand", ec.treesitPatternExpand_autogen, 1)
	ec.defSubr5(nil, "treesit-query-capture", ec.treesitQueryCapture_autogen, 2)
	ec.defSubr3(nil, "treesit-query-compile", ec.treesitQueryCompile_autogen, 2)
	ec.defSubr1(nil, "treesit-query-expand", ec.treesitQueryExpand_autogen, 1)
	ec.defSubr1(nil, "treesit-query-language", ec.treesitQueryLanguage_autogen, 1)
	ec.defSubr1(nil, "treesit-query-p", ec.treesitQueryP_autogen, 1)
	ec.defSubr4(nil, "treesit-search-forward", ec.treesitSearchForward_autogen, 2)
	ec.defSubr5(nil, "treesit-search-subtree", ec.treesitSearchSubtree_autogen, 2)
	ec.defSubr1(nil, "treesit-subtree-stat", ec.treesitSubtreeStat_autogen, 1)
	ec.defSubr0(nil, "undo-boundary", ec.undoBoundary_autogen)
	ec.defSubr1(nil, "w16-get-clipboard-data", ec.w16GetClipboardData_autogen, 0)
	ec.defSubr2(nil, "w16-selection-exists-p", ec.w16SelectionExistsP_autogen, 0)
	ec.defSubr2(nil, "w16-set-clipboard-data", ec.w16SetClipboardData_autogen, 1)
	ec.defSubr0(nil, "get-screen-color", ec.getScreenColor_autogen)
	ec.defSubr1(nil, "set-cursor-size", ec.setCursorSize_autogen, 1)
	ec.defSubr2(nil, "set-screen-color", ec.setScreenColor_autogen, 2)
	ec.defSubr0(nil, "w32-battery-status", ec.w32BatteryStatus_autogen)
	ec.defSubr0(nil, "default-printer-name", ec.defaultPrinterName_autogen)
	ec.defSubr1(nil, "set-message-beep", ec.setMessageBeep_autogen, 1)
	ec.defSubr1(nil, "system-move-file-to-trash", ec.systemMoveFileToTrash_autogen, 1)
	ec.defSubr0(nil, "w32--menu-bar-in-use", ec.w32MenuBarInUse_autogen)
	ec.defSubr4(nil, "w32-define-rgb-color", ec.w32DefineRgbColor_autogen, 4)
	ec.defSubr1(nil, "w32-display-monitor-attributes-list", ec.w32DisplayMonitorAttributesList_autogen, 0)
	ec.defSubr2(nil, "w32-frame-edges", ec.w32FrameEdges_autogen, 0)
	ec.defSubr1(nil, "w32-frame-geometry", ec.w32FrameGeometry_autogen, 0)
	ec.defSubr1(nil, "w32-frame-list-z-order", ec.w32FrameListZOrder_autogen, 0)
	ec.defSubr3(nil, "w32-frame-restack", ec.w32FrameRestack_autogen, 2)
	ec.defSubr0(nil, "w32-get-ime-open-status", ec.w32GetImeOpenStatus_autogen)
	ec.defSubr0(nil, "w32-mouse-absolute-pixel-position", ec.w32MouseAbsolutePixelPosition_autogen)
	ec.defSubr1(nil, "w32-notification-close", ec.w32NotificationClose_autogen, 1)
	ec.defSubrM(nil, "w32-notification-notify", ec.w32NotificationNotify_autogen, 0)
	ec.defSubr3(nil, "w32-read-registry", ec.w32ReadRegistry_autogen, 3)
	ec.defSubr1(nil, "w32-reconstruct-hot-key", ec.w32ReconstructHotKey_autogen, 1)
	ec.defSubr1(nil, "w32-register-hot-key", ec.w32RegisterHotKey_autogen, 1)
	ec.defSubr0(nil, "w32-registered-hot-keys", ec.w32RegisteredHotKeys_autogen)
	ec.defSubr2(nil, "w32-send-sys-command", ec.w32SendSysCommand_autogen, 1)
	ec.defSubr1(nil, "w32-set-ime-open-status", ec.w32SetImeOpenStatus_autogen, 1)
	ec.defSubr2(nil, "w32-set-mouse-absolute-pixel-position", ec.w32SetMouseAbsolutePixelPosition_autogen, 2)
	ec.defSubr1(nil, "w32-set-wallpaper", ec.w32SetWallpaper_autogen, 1)
	ec.defSubr4(nil, "w32-shell-execute", ec.w32ShellExecute_autogen, 2)
	ec.defSubr2(nil, "w32-toggle-lock-key", ec.w32ToggleLockKey_autogen, 1)
	ec.defSubr1(nil, "w32-unregister-hot-key", ec.w32UnregisterHotKey_autogen, 1)
	ec.defSubr2(nil, "w32-window-exists-p", ec.w32WindowExistsP_autogen, 2)
	ec.defSubr6(nil, "x-change-window-property", ec.xChangeWindowProperty_autogen, 2)
	ec.defSubr7(nil, "x-change-window-property", ec.xChangeWindowProperty_1_autogen, 2)
	ec.defSubr2(nil, "x-delete-window-property", ec.xDeleteWindowProperty_autogen, 1)
	ec.defSubr3(nil, "x-delete-window-property", ec.xDeleteWindowProperty_1_autogen, 1)
	ec.defSubr2(nil, "x-synchronize", ec.xSynchronize_autogen, 1)
	ec.defSubr2(nil, "x-synchronize", ec.xSynchronize_1_autogen, 1)
	ec.defSubr6(nil, "x-window-property", ec.xWindowProperty_autogen, 1)
	ec.defSubr6(nil, "x-window-property", ec.xWindowProperty_1_autogen, 1)
	ec.defSubr3(nil, "w32notify-add-watch", ec.w32notifyAddWatch_autogen, 3)
	ec.defSubr1(nil, "w32notify-rm-watch", ec.w32notifyRmWatch_autogen, 1)
	ec.defSubr1(nil, "w32notify-valid-p", ec.w32notifyValidP_autogen, 1)
	ec.defSubr1(nil, "w32-application-type", ec.w32ApplicationType_autogen, 1)
	ec.defSubr1(nil, "w32-get-codepage-charset", ec.w32GetCodepageCharset_autogen, 1)
	ec.defSubr0(nil, "w32-get-console-codepage", ec.w32GetConsoleCodepage_autogen)
	ec.defSubr0(nil, "w32-get-console-output-codepage", ec.w32GetConsoleOutputCodepage_autogen)
	ec.defSubr0(nil, "w32-get-current-locale-id", ec.w32GetCurrentLocaleId_autogen)
	ec.defSubr1(nil, "w32-get-default-locale-id", ec.w32GetDefaultLocaleId_autogen, 0)
	ec.defSubr0(nil, "w32-get-keyboard-layout", ec.w32GetKeyboardLayout_autogen)
	ec.defSubr2(nil, "w32-get-locale-info", ec.w32GetLocaleInfo_autogen, 1)
	ec.defSubr0(nil, "w32-get-valid-codepages", ec.w32GetValidCodepages_autogen)
	ec.defSubr0(nil, "w32-get-valid-keyboard-layouts", ec.w32GetValidKeyboardLayouts_autogen)
	ec.defSubr0(nil, "w32-get-valid-locale-ids", ec.w32GetValidLocaleIds_autogen)
	ec.defSubr1(nil, "w32-has-winsock", ec.w32HasWinsock_autogen, 0)
	ec.defSubr1(nil, "w32-long-file-name", ec.w32LongFileName_autogen, 1)
	ec.defSubr1(nil, "w32-set-console-codepage", ec.w32SetConsoleCodepage_autogen, 1)
	ec.defSubr1(nil, "w32-set-console-output-codepage", ec.w32SetConsoleOutputCodepage_autogen, 1)
	ec.defSubr1(nil, "w32-set-current-locale", ec.w32SetCurrentLocale_autogen, 1)
	ec.defSubr1(nil, "w32-set-keyboard-layout", ec.w32SetKeyboardLayout_autogen, 1)
	ec.defSubr2(nil, "w32-set-process-priority", ec.w32SetProcessPriority_autogen, 2)
	ec.defSubr1(nil, "w32-short-file-name", ec.w32ShortFileName_autogen, 1)
	ec.defSubr0(nil, "w32-unload-winsock", ec.w32UnloadWinsock_autogen)
	ec.defSubr1(nil, "w32-get-clipboard-data", ec.w32GetClipboardData_autogen, 0)
	ec.defSubr2(nil, "w32-selection-exists-p", ec.w32SelectionExistsP_autogen, 0)
	ec.defSubr2(nil, "w32-selection-targets", ec.w32SelectionTargets_autogen, 0)
	ec.defSubr2(nil, "w32-set-clipboard-data", ec.w32SetClipboardData_autogen, 1)
	ec.defSubr2(nil, "coordinates-in-window-p", ec.coordinatesInWindowP_autogen, 2)
	ec.defSubr1(nil, "current-window-configuration", ec.currentWindowConfiguration_autogen, 0)
	ec.defSubr2(nil, "delete-other-windows-internal", ec.deleteOtherWindowsInternal_autogen, 0)
	ec.defSubr1(nil, "delete-window-internal", ec.deleteWindowInternal_autogen, 1)
	ec.defSubr1(nil, "force-window-update", ec.forceWindowUpdate_autogen, 0)
	ec.defSubr1(nil, "frame-first-window", ec.frameFirstWindow_autogen, 0)
	ec.defSubr1(nil, "frame-old-selected-window", ec.frameOldSelectedWindow_autogen, 0)
	ec.defSubr1(nil, "frame-root-window", ec.frameRootWindow_autogen, 0)
	ec.defSubr1(nil, "frame-selected-window", ec.frameSelectedWindow_autogen, 0)
	ec.defSubr2(nil, "get-buffer-window", ec.getBufferWindow_autogen, 0)
	ec.defSubr0(nil, "minibuffer-selected-window", ec.minibufferSelectedWindow_autogen)
	ec.defSubr1(nil, "minibuffer-window", ec.minibufferWindow_autogen, 0)
	ec.defSubr1(nil, "move-to-window-line", ec.moveToWindowLine_autogen, 1)
	ec.defSubr3(nil, "next-window", ec.nextWindow_autogen, 0)
	ec.defSubr0(nil, "old-selected-window", ec.oldSelectedWindow_autogen)
	ec.defSubr0(nil, "other-window-for-scrolling", ec.otherWindowForScrolling_autogen)
	ec.defSubr3(nil, "pos-visible-in-window-p", ec.posVisibleInWindowP_autogen, 0)
	ec.defSubr3(nil, "previous-window", ec.previousWindow_autogen, 0)
	ec.defSubr2(nil, "recenter", ec.recenter_autogen, 0)
	ec.defSubr1(nil, "resize-mini-window-internal", ec.resizeMiniWindowInternal_autogen, 1)
	ec.defSubr1(nil, "run-window-configuration-change-hook", ec.runWindowConfigurationChangeHook_autogen, 0)
	ec.defSubr1(nil, "run-window-scroll-functions", ec.runWindowScrollFunctions_autogen, 0)
	ec.defSubr1(nil, "scroll-down", ec.scrollDown_autogen, 0)
	ec.defSubr2(nil, "scroll-left", ec.scrollLeft_autogen, 0)
	ec.defSubr2(nil, "scroll-right", ec.scrollRight_autogen, 0)
	ec.defSubr1(nil, "scroll-up", ec.scrollUp_autogen, 0)
	ec.defSubr2(nil, "select-window", ec.selectWindow_autogen, 1)
	ec.defSubr0(nil, "selected-window", ec.selectedWindow_autogen)
	ec.defSubr3(nil, "set-frame-selected-window", ec.setFrameSelectedWindow_autogen, 2)
	ec.defSubr3(nil, "set-window-buffer", ec.setWindowBuffer_autogen, 2)
	ec.defSubr2(nil, "set-window-combination-limit", ec.setWindowCombinationLimit_autogen, 2)
	ec.defSubr3(nil, "set-window-configuration", ec.setWindowConfiguration_autogen, 1)
	ec.defSubr2(nil, "set-window-dedicated-p", ec.setWindowDedicatedP_autogen, 2)
	ec.defSubr2(nil, "set-window-display-table", ec.setWindowDisplayTable_autogen, 2)
	ec.defSubr5(nil, "set-window-fringes", ec.setWindowFringes_autogen, 2)
	ec.defSubr2(nil, "set-window-hscroll", ec.setWindowHscroll_autogen, 2)
	ec.defSubr3(nil, "set-window-margins", ec.setWindowMargins_autogen, 2)
	ec.defSubr2(nil, "set-window-new-normal", ec.setWindowNewNormal_autogen, 1)
	ec.defSubr3(nil, "set-window-new-pixel", ec.setWindowNewPixel_autogen, 2)
	ec.defSubr3(nil, "set-window-new-total", ec.setWindowNewTotal_autogen, 2)
	ec.defSubr2(nil, "set-window-next-buffers", ec.setWindowNextBuffers_autogen, 2)
	ec.defSubr3(nil, "set-window-parameter", ec.setWindowParameter_autogen, 3)
	ec.defSubr2(nil, "set-window-point", ec.setWindowPoint_autogen, 2)
	ec.defSubr2(nil, "set-window-prev-buffers", ec.setWindowPrevBuffers_autogen, 2)
	ec.defSubr6(nil, "set-window-scroll-bars", ec.setWindowScrollBars_autogen, 1)
	ec.defSubr3(nil, "set-window-start", ec.setWindowStart_autogen, 2)
	ec.defSubr4(nil, "set-window-vscroll", ec.setWindowVscroll_autogen, 2)
	ec.defSubr4(nil, "split-window-internal", ec.splitWindowInternal_autogen, 4)
	ec.defSubr3(nil, "window-at", ec.windowAt_autogen, 2)
	ec.defSubr2(nil, "window-body-height", ec.windowBodyHeight_autogen, 0)
	ec.defSubr2(nil, "window-body-width", ec.windowBodyWidth_autogen, 0)
	ec.defSubr1(nil, "window-bottom-divider-width", ec.windowBottomDividerWidth_autogen, 0)
	ec.defSubr1(nil, "window-buffer", ec.windowBuffer_autogen, 0)
	ec.defSubr1(nil, "window-bump-use-time", ec.windowBumpUseTime_autogen, 0)
	ec.defSubr1(nil, "window-combination-limit", ec.windowCombinationLimit_autogen, 1)
	ec.defSubr2(nil, "window-configuration-equal-p", ec.windowConfigurationEqualP_autogen, 2)
	ec.defSubr1(nil, "window-configuration-frame", ec.windowConfigurationFrame_autogen, 1)
	ec.defSubr1(nil, "window-configuration-p", ec.windowConfigurationP_autogen, 1)
	ec.defSubr1(nil, "window-dedicated-p", ec.windowDedicatedP_autogen, 0)
	ec.defSubr1(nil, "window-display-table", ec.windowDisplayTable_autogen, 0)
	ec.defSubr2(nil, "window-end", ec.windowEnd_autogen, 0)
	ec.defSubr1(nil, "window-frame", ec.windowFrame_autogen, 0)
	ec.defSubr1(nil, "window-fringes", ec.windowFringes_autogen, 0)
	ec.defSubr1(nil, "window-header-line-height", ec.windowHeaderLineHeight_autogen, 0)
	ec.defSubr1(nil, "window-hscroll", ec.windowHscroll_autogen, 0)
	ec.defSubr1(nil, "window-left-child", ec.windowLeftChild_autogen, 0)
	ec.defSubr1(nil, "window-left-column", ec.windowLeftColumn_autogen, 0)
	ec.defSubr2(nil, "window-line-height", ec.windowLineHeight_autogen, 0)
	ec.defSubr6(nil, "window-lines-pixel-dimensions", ec.windowLinesPixelDimensions_autogen, 0)
	ec.defSubr3(nil, "window-list", ec.windowList_autogen, 0)
	ec.defSubr3(nil, "window-list-1", ec.windowList1_autogen, 0)
	ec.defSubr1(nil, "window-live-p", ec.windowLiveP_autogen, 1)
	ec.defSubr1(nil, "window-margins", ec.windowMargins_autogen, 0)
	ec.defSubr1(nil, "window-minibuffer-p", ec.windowMinibufferP_autogen, 0)
	ec.defSubr1(nil, "window-mode-line-height", ec.windowModeLineHeight_autogen, 0)
	ec.defSubr1(nil, "window-new-normal", ec.windowNewNormal_autogen, 0)
	ec.defSubr1(nil, "window-new-pixel", ec.windowNewPixel_autogen, 0)
	ec.defSubr1(nil, "window-new-total", ec.windowNewTotal_autogen, 0)
	ec.defSubr1(nil, "window-next-buffers", ec.windowNextBuffers_autogen, 0)
	ec.defSubr1(nil, "window-next-sibling", ec.windowNextSibling_autogen, 0)
	ec.defSubr2(nil, "window-normal-size", ec.windowNormalSize_autogen, 0)
	ec.defSubr1(nil, "window-old-body-pixel-height", ec.windowOldBodyPixelHeight_autogen, 0)
	ec.defSubr1(nil, "window-old-body-pixel-width", ec.windowOldBodyPixelWidth_autogen, 0)
	ec.defSubr1(nil, "window-old-buffer", ec.windowOldBuffer_autogen, 0)
	ec.defSubr1(nil, "window-old-pixel-height", ec.windowOldPixelHeight_autogen, 0)
	ec.defSubr1(nil, "window-old-pixel-width", ec.windowOldPixelWidth_autogen, 0)
	ec.defSubr1(nil, "window-old-point", ec.windowOldPoint_autogen, 0)
	ec.defSubr2(nil, "window-parameter", ec.windowParameter_autogen, 2)
	ec.defSubr1(nil, "window-parameters", ec.windowParameters_autogen, 0)
	ec.defSubr1(nil, "window-parent", ec.windowParent_autogen, 0)
	ec.defSubr1(nil, "window-pixel-height", ec.windowPixelHeight_autogen, 0)
	ec.defSubr1(nil, "window-pixel-left", ec.windowPixelLeft_autogen, 0)
	ec.defSubr1(nil, "window-pixel-top", ec.windowPixelTop_autogen, 0)
	ec.defSubr1(nil, "window-pixel-width", ec.windowPixelWidth_autogen, 0)
	ec.defSubr1(nil, "window-point", ec.windowPoint_autogen, 0)
	ec.defSubr1(nil, "window-prev-buffers", ec.windowPrevBuffers_autogen, 0)
	ec.defSubr1(nil, "window-prev-sibling", ec.windowPrevSibling_autogen, 0)
	ec.defSubr2(nil, "window-resize-apply", ec.windowResizeApply_autogen, 0)
	ec.defSubr2(nil, "window-resize-apply-total", ec.windowResizeApplyTotal_autogen, 0)
	ec.defSubr1(nil, "window-right-divider-width", ec.windowRightDividerWidth_autogen, 0)
	ec.defSubr1(nil, "window-scroll-bar-height", ec.windowScrollBarHeight_autogen, 0)
	ec.defSubr1(nil, "window-scroll-bar-width", ec.windowScrollBarWidth_autogen, 0)
	ec.defSubr1(nil, "window-scroll-bars", ec.windowScrollBars_autogen, 0)
	ec.defSubr1(nil, "window-start", ec.windowStart_autogen, 0)
	ec.defSubr1(nil, "window-tab-line-height", ec.windowTabLineHeight_autogen, 0)
	ec.defSubr2(nil, "window-text-height", ec.windowTextHeight_autogen, 0)
	ec.defSubr2(nil, "window-text-width", ec.windowTextWidth_autogen, 0)
	ec.defSubr1(nil, "window-top-child", ec.windowTopChild_autogen, 0)
	ec.defSubr1(nil, "window-top-line", ec.windowTopLine_autogen, 0)
	ec.defSubr2(nil, "window-total-height", ec.windowTotalHeight_autogen, 0)
	ec.defSubr2(nil, "window-total-width", ec.windowTotalWidth_autogen, 0)
	ec.defSubr1(nil, "window-use-time", ec.windowUseTime_autogen, 0)
	ec.defSubr1(nil, "window-valid-p", ec.windowValidP_autogen, 1)
	ec.defSubr2(nil, "window-vscroll", ec.windowVscroll_autogen, 0)
	ec.defSubr1(nil, "windowp", ec.windowp_autogen, 1)
	ec.defSubr4(nil, "bidi-find-overridden-directionality", ec.bidiFindOverriddenDirectionality_autogen, 3)
	ec.defSubr1(nil, "bidi-resolved-levels", ec.bidiResolvedLevels_autogen, 0)
	ec.defSubr4(nil, "buffer-text-pixel-size", ec.bufferTextPixelSize_autogen, 0)
	ec.defSubr1(nil, "current-bidi-paragraph-direction", ec.currentBidiParagraphDirection_autogen, 0)
	ec.defSubr0(nil, "display--line-is-continued-p", ec.displayLineIsContinuedP_autogen)
	ec.defSubr0(nil, "dump-frame-glyph-matrix", ec.dumpFrameGlyphMatrix_autogen)
	ec.defSubr1(nil, "dump-glyph-matrix", ec.dumpGlyphMatrix_autogen, 0)
	ec.defSubr2(nil, "dump-glyph-row", ec.dumpGlyphRow_autogen, 1)
	ec.defSubr2(nil, "dump-tab-bar-row", ec.dumpTabBarRow_autogen, 1)
	ec.defSubr2(nil, "dump-tool-bar-row", ec.dumpToolBarRow_autogen, 1)
	ec.defSubr4(nil, "format-mode-line", ec.formatModeLine_autogen, 1)
	ec.defSubr4(nil, "get-display-property", ec.getDisplayProperty_autogen, 2)
	ec.defSubr1(nil, "invisible-p", ec.invisibleP_autogen, 1)
	ec.defSubr0(nil, "line-pixel-height", ec.linePixelHeight_autogen)
	ec.defSubr0(nil, "long-line-optimizations-p", ec.longLineOptimizationsP_autogen)
	ec.defSubr3(nil, "lookup-image-map", ec.lookupImageMap_autogen, 3)
	ec.defSubr1(nil, "move-point-visually", ec.movePointVisually_autogen, 1)
	ec.defSubr4(nil, "set-buffer-redisplay", ec.setBufferRedisplay_autogen, 4)
	ec.defSubr2(nil, "tab-bar-height", ec.tabBarHeight_autogen, 0)
	ec.defSubr2(nil, "tool-bar-height", ec.toolBarHeight_autogen, 0)
	ec.defSubr1(nil, "trace-redisplay", ec.traceRedisplay_autogen, 0)
	ec.defSubrM(nil, "trace-to-stderr", ec.traceToStderr_autogen, 1)
	ec.defSubr7(nil, "window-text-pixel-size", ec.windowTextPixelSize_autogen, 0)
	ec.defSubr1(nil, "bitmap-spec-p", ec.bitmapSpecP_autogen, 1)
	ec.defSubr1(nil, "clear-face-cache", ec.clearFaceCache_autogen, 0)
	ec.defSubr4(nil, "color-distance", ec.colorDistance_autogen, 2)
	ec.defSubr2(nil, "color-gray-p", ec.colorGrayP_autogen, 1)
	ec.defSubr3(nil, "color-supported-p", ec.colorSupportedP_autogen, 1)
	ec.defSubr1(nil, "color-values-from-color-spec", ec.colorValuesFromColorSpec_autogen, 1)
	ec.defSubr2(nil, "display-supports-face-attributes-p", ec.displaySupportsFaceAttributesP_autogen, 1)
	ec.defSubr0(nil, "dump-colors", ec.dumpColors_autogen)
	ec.defSubr1(nil, "dump-face", ec.dumpFace_autogen, 0)
	ec.defSubr2(nil, "face-attribute-relative-p", ec.faceAttributeRelativeP_autogen, 2)
	ec.defSubr1(nil, "face-attributes-as-vector", ec.faceAttributesAsVector_autogen, 1)
	ec.defSubr3(nil, "face-font", ec.faceFont_autogen, 1)
	ec.defSubr1(nil, "frame--face-hash-table", ec.frameFaceHashTable_autogen, 0)
	ec.defSubr4(nil, "internal-copy-lisp-face", ec.internalCopyLispFace_autogen, 4)
	ec.defSubr3(nil, "internal-face-x-get-resource", ec.internalFaceXGetResource_autogen, 2)
	ec.defSubr3(nil, "internal-get-lisp-face-attribute", ec.internalGetLispFaceAttribute_autogen, 2)
	ec.defSubr1(nil, "internal-lisp-face-attribute-values", ec.internalLispFaceAttributeValues_autogen, 1)
	ec.defSubr2(nil, "internal-lisp-face-empty-p", ec.internalLispFaceEmptyP_autogen, 1)
	ec.defSubr3(nil, "internal-lisp-face-equal-p", ec.internalLispFaceEqualP_autogen, 2)
	ec.defSubr2(nil, "internal-lisp-face-p", ec.internalLispFaceP_autogen, 1)
	ec.defSubr2(nil, "internal-make-lisp-face", ec.internalMakeLispFace_autogen, 1)
	ec.defSubr2(nil, "internal-merge-in-global-face", ec.internalMergeInGlobalFace_autogen, 2)
	ec.defSubr1(nil, "internal-set-alternative-font-family-alist", ec.internalSetAlternativeFontFamilyAlist_autogen, 1)
	ec.defSubr1(nil, "internal-set-alternative-font-registry-alist", ec.internalSetAlternativeFontRegistryAlist_autogen, 1)
	ec.defSubr1(nil, "internal-set-font-selection-order", ec.internalSetFontSelectionOrder_autogen, 1)
	ec.defSubr4(nil, "internal-set-lisp-face-attribute", ec.internalSetLispFaceAttribute_autogen, 3)
	ec.defSubr4(nil, "internal-set-lisp-face-attribute-from-resource", ec.internalSetLispFaceAttributeFromResource_autogen, 3)
	ec.defSubr3(nil, "merge-face-attribute", ec.mergeFaceAttribute_autogen, 3)
	ec.defSubr0(nil, "show-face-resources", ec.showFaceResources_autogen)
	ec.defSubr1(nil, "tty-suppress-bold-inverse-default-colors", ec.ttySuppressBoldInverseDefaultColors_autogen, 1)
	ec.defSubr2(nil, "x-family-fonts", ec.xFamilyFonts_autogen, 0)
	ec.defSubr5(nil, "x-list-fonts", ec.xListFonts_autogen, 1)
	ec.defSubr1(nil, "x-load-color-file", ec.xLoadColorFile_autogen, 1)
	ec.defSubr1(nil, "x-backspace-delete-keys-p", ec.xBackspaceDeleteKeysP_autogen, 0)
	ec.defSubr6(nil, "x-begin-drag", ec.xBeginDrag_autogen, 1)
	ec.defSubr1(nil, "x-display-monitor-attributes-list", ec.xDisplayMonitorAttributesList_autogen, 0)
	ec.defSubr2(nil, "x-display-set-last-user-time", ec.xDisplayLastUserTime_autogen, 1)
	ec.defSubr2(nil, "x-frame-edges", ec.xFrameEdges_autogen, 0)
	ec.defSubr1(nil, "x-frame-geometry", ec.xFrameGeometry_autogen, 0)
	ec.defSubr1(nil, "x-frame-list-z-order", ec.xFrameListZOrder_autogen, 0)
	ec.defSubr3(nil, "x-frame-restack", ec.xFrameRestack_autogen, 2)
	ec.defSubr1(nil, "x-get-modifier-masks", ec.xGetModifierMasks_autogen, 0)
	ec.defSubr0(nil, "x-get-page-setup", ec.xGetPageSetup_autogen)
	ec.defSubr1(nil, "x-internal-focus-input-context", ec.xInternalFocusInputContext_autogen, 1)
	ec.defSubr0(nil, "x-mouse-absolute-pixel-position", ec.xMouseAbsolutePixelPosition_autogen)
	ec.defSubr0(nil, "x-page-setup-dialog", ec.xPageSetupDialog_autogen)
	ec.defSubr1(nil, "x-print-frames-dialog", ec.xPrintFramesDialog_autogen, 0)
	ec.defSubr1(nil, "x-server-input-extension-version", ec.xServerInputExtensionVersion_autogen, 0)
	ec.defSubr2(nil, "x-set-mouse-absolute-pixel-position", ec.xSetMouseAbsolutePixelPosition_autogen, 2)
	ec.defSubr6(nil, "x-translate-coordinates", ec.xTranslateCoordinates_autogen, 1)
	ec.defSubr0(nil, "x-uses-old-gtk-dialog", ec.xUsesOldGtkDialog_autogen)
	ec.defSubr3(nil, "x-window-property-attributes", ec.xWindowPropertyAttributes_autogen, 1)
	ec.defSubr1(nil, "x-wm-set-size-hint", ec.xWmSetSizeHint_autogen, 0)
	ec.defSubr0(nil, "libxml-available-p", ec.libxmlAvailableP_autogen)
	ec.defSubr4(nil, "libxml-parse-html-region", ec.libxmlParseHtmlRegion_autogen, 0)
	ec.defSubr4(nil, "libxml-parse-xml-region", ec.libxmlParseXmlRegion_autogen, 0)
	ec.defSubr3(nil, "x-disown-selection-internal", ec.xDisownSelectionInternal_autogen, 1)
	ec.defSubr2(nil, "x-get-atom-name", ec.xGetAtomName_autogen, 1)
	ec.defSubr2(nil, "x-get-local-selection", ec.xGetLocalSelection_autogen, 0)
	ec.defSubr4(nil, "x-get-selection-internal", ec.xGetSelectionInternal_autogen, 2)
	ec.defSubr3(nil, "x-own-selection-internal", ec.xOwnSelectionInternal_autogen, 2)
	ec.defSubr2(nil, "x-register-dnd-atom", ec.xRegisterDndAtom_autogen, 1)
	ec.defSubr2(nil, "x-selection-exists-p", ec.xSelectionExistsP_autogen, 0)
	ec.defSubr2(nil, "x-selection-owner-p", ec.xSelectionOwnerP_autogen, 0)
	ec.defSubr6(nil, "x-send-client-message", ec.xSendClientMessage_autogen, 6)
	ec.defSubr0(nil, "tool-bar-get-system-style", ec.toolBarGetSystemStyle_autogen)
	ec.defSubr1(nil, "handle-save-session", ec.handleSaveSession_autogen, 1)
	ec.defSubr1(nil, "delete-xwidget-view", ec.deleteXwidgetView_autogen, 1)
	ec.defSubr1(nil, "get-buffer-xwidgets", ec.getBufferXwidgets_autogen, 1)
	ec.defSubr1(nil, "kill-xwidget", ec.killXwidget_autogen, 1)
	ec.defSubr7(nil, "make-xwidget", ec.makeXwidget_autogen, 4)
	ec.defSubr2(nil, "set-xwidget-buffer", ec.setXwidgetBuffer_autogen, 2)
	ec.defSubr2(nil, "set-xwidget-plist", ec.setXwidgetPlist_autogen, 2)
	ec.defSubr2(nil, "set-xwidget-query-on-exit-flag", ec.setXwidgetQueryOnExitFlag_autogen, 2)
	ec.defSubr1(nil, "xwidget-buffer", ec.xwidgetBuffer_autogen, 1)
	ec.defSubr1(nil, "xwidget-info", ec.xwidgetInfo_autogen, 1)
	ec.defSubr1(nil, "xwidget-live-p", ec.xwidgetLiveP_autogen, 1)
	ec.defSubr3(nil, "xwidget-perform-lispy-event", ec.xwidgetPerformLispyEvent_autogen, 2)
	ec.defSubr1(nil, "xwidget-plist", ec.xwidgetPlist_autogen, 1)
	ec.defSubr1(nil, "xwidget-query-on-exit-flag", ec.xwidgetQueryOnExitFlag_autogen, 1)
	ec.defSubr3(nil, "xwidget-resize", ec.xwidgetResize_autogen, 3)
	ec.defSubr1(nil, "xwidget-size-request", ec.xwidgetSizeRequest_autogen, 1)
	ec.defSubr1(nil, "xwidget-view-info", ec.xwidgetViewInfo_autogen, 1)
	ec.defSubr2(nil, "xwidget-view-lookup", ec.xwidgetViewLookup_autogen, 1)
	ec.defSubr1(nil, "xwidget-view-model", ec.xwidgetViewModel_autogen, 1)
	ec.defSubr1(nil, "xwidget-view-p", ec.xwidgetViewP_autogen, 1)
	ec.defSubr1(nil, "xwidget-view-window", ec.xwidgetViewWindow_autogen, 1)
	ec.defSubr2(nil, "xwidget-webkit-back-forward-list", ec.xwidgetWebkitBackForwardList_autogen, 1)
	ec.defSubr1(nil, "xwidget-webkit-estimated-load-progress", ec.xwidgetWebkitEstimatedLoadProgress_autogen, 1)
	ec.defSubr3(nil, "xwidget-webkit-execute-script", ec.xwidgetWebkitExecuteScript_autogen, 2)
	ec.defSubr1(nil, "xwidget-webkit-finish-search", ec.xwidgetWebkitFinishSearch_autogen, 1)
	ec.defSubr2(nil, "xwidget-webkit-goto-history", ec.xwidgetWebkitGotoHistory_autogen, 2)
	ec.defSubr2(nil, "xwidget-webkit-goto-uri", ec.xwidgetWebkitGotoUri_autogen, 2)
	ec.defSubr3(nil, "xwidget-webkit-load-html", ec.xwidgetWebkitLoadHtml_autogen, 2)
	ec.defSubr1(nil, "xwidget-webkit-next-result", ec.xwidgetWebkitNextResult_autogen, 1)
	ec.defSubr1(nil, "xwidget-webkit-previous-result", ec.xwidgetWebkitPreviousResult_autogen, 1)
	ec.defSubr5(nil, "xwidget-webkit-search", ec.xwidgetWebkitSearch_autogen, 2)
	ec.defSubr2(nil, "xwidget-webkit-set-cookie-storage-file", ec.xwidgetWebkitSetCookieStorageFile_autogen, 2)
	ec.defSubr1(nil, "xwidget-webkit-stop-loading", ec.xwidgetWebkitStopLoading_autogen, 1)
	ec.defSubr1(nil, "xwidget-webkit-title", ec.xwidgetWebkitTitle_autogen, 1)
	ec.defSubr1(nil, "xwidget-webkit-uri", ec.xwidgetWebkitUri_autogen, 1)
	ec.defSubr2(nil, "xwidget-webkit-zoom", ec.xwidgetWebkitZoom_autogen, 2)
	ec.defSubr1(nil, "xwidgetp", ec.xwidgetp_autogen, 1)
}

// Subroutines count: 1716
// Constants count: 2
